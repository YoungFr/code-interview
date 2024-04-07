用 Go 语言实现的 KV 存储项目 [`diskv`](https://github.com/peterbourgon/diskv) 的源代码分析：

```go
package diskv

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	"github.com/google/btree"
)

const (
	defaultBasePath             = "diskv" // 键值对默认存放的基目录
	defaultFilePerm os.FileMode = 0666    // 文件默认是读、写权限
	defaultPathPerm os.FileMode = 0777    // 目录默认是读、写和执行权限
)

var (
	errCanceled        = errors.New("canceled")
	errEmptyKey        = errors.New("empty key")
	errBadKey          = errors.New("bad key")
	errImportDirectory = errors.New("can't import a directory")
)

type PathKey struct {
	Path        []string
	FileName    string // 键对应的值存储在 <basedir>/Path[0]/Path[1]/.../Path[n-1]/FileName 文件中
	originalKey string // 用于存储键
}

// 函数 TransformFunction 将键转换成字符串切片，切片中的每个元素都表示一个目录
// 比如将键 "abcdef" 转换成 ["ab", "cde", "f"] 切片，那么值最终存储在 <basedir>/ab/cde/f/abcdef 文件中
// 最简单的转换函数返回一个长度为 0 的字符串切片，此时值存储在 <basedir>/abcdef 文件中
type TransformFunction func(key string) []string

// 将键转换成字符串切片并和文件名一起存放在 PathKey 结构体中
type AdvancedTransformFunction func(key string) *PathKey

// 如果提供了 AdvancedTransform 函数就必须同时提供 InverseTransform 函数
type InverseTransformFunction func(pathKey *PathKey) string

var (
	defaultAdvancedTransform = func(key string) *PathKey { return &PathKey{Path: []string{}, FileName: key} }
	defaultInverseTransform  = func(pathKey *PathKey) string { return pathKey.FileName }
)

// 结构体 Options 所有的字段都是可选的
type Options struct {
	BasePath          string
	Transform         TransformFunction
	AdvancedTransform AdvancedTransformFunction
	InverseTransform  InverseTransformFunction
	CacheSizeMax      uint64       // 最大缓存的大小，单位是字节
	PathPerm          os.FileMode  // 路径权限
	FilePerm          os.FileMode  // 文件权限
	TempDir           string       // TempDir MUST be on the same device/partition as BasePath.
	Index             Index        // 用于实现键的排序
	IndexLess         LessFunction // 用于实现键的排序
	Compression       Compression  // 用于实现数据压缩
}

type Diskv struct {
	Options
	mu        sync.RWMutex
	cache     map[string][]byte
	cacheSize uint64
}

// 将 TransformFunction 类型转换为等价的 AdvancedTransformFunction 类型
func convertToAdvancedTransform(oldFunc func(s string) []string) AdvancedTransformFunction {
	return func(s string) *PathKey { return &PathKey{Path: oldFunc(s), FileName: s} }
}

func New(o Options) *Diskv {
	if o.BasePath == "" {
		o.BasePath = defaultBasePath
	}
	// 设置转换函数，最终要设置的是 AdvancedTransform 和 InverseTransform 两个字段
	if o.AdvancedTransform == nil {
		if o.Transform == nil {
			o.AdvancedTransform = defaultAdvancedTransform
		} else {
			o.AdvancedTransform = convertToAdvancedTransform(o.Transform)
		}
		if o.InverseTransform == nil {
			o.InverseTransform = defaultInverseTransform
		}
	} else {
		if o.InverseTransform == nil {
			panic("You must provide an InverseTransform function in advanced mode")
		}
	}
	if o.PathPerm == 0 {
		o.PathPerm = defaultPathPerm
	}
	if o.FilePerm == 0 {
		o.FilePerm = defaultFilePerm
	}
	d := &Diskv{
		Options:   o,
		cache:     map[string][]byte{},
		cacheSize: 0,
	}
	if d.Index != nil && d.IndexLess != nil {
		d.Index.Initialize(d.IndexLess, d.Keys(nil))
	}
	return d
}

func (d *Diskv) WriteString(key string, val string) error {
	return d.Write(key, []byte(val))
}

// 同步地将 key-value 对写入磁盘文件
func (d *Diskv) Write(key string, val []byte) error {
	return d.WriteStream(key, bytes.NewReader(val), false)
}

// 函数 transform 是对 AdvancedTransform 的包装，目的是为了设置 PathKey 私有的 originalKey 字段
func (d *Diskv) transform(key string) (pathKey *PathKey) {
	pathKey = d.AdvancedTransform(key)
	pathKey.originalKey = key
	return pathKey
}

func (d *Diskv) WriteStream(key string, r io.Reader, sync bool) error {
	if len(key) <= 0 {
		return errEmptyKey
	}
	// 检查文件路径的合法性，即每一部分都不能包含操作系统规定的路径分隔符
	pathKey := d.transform(key)
	for _, pathPart := range pathKey.Path {
		if strings.ContainsRune(pathPart, os.PathSeparator) {
			return errBadKey
		}
	}
	if strings.ContainsRune(pathKey.FileName, os.PathSeparator) {
		return errBadKey
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.writeStreamWithLock(pathKey, r, sync)
}

// 给定一个 PathKey 返回它所表示的目录名
// 也就是 <basedir>/Path[0]/Path[1]/.../Path[n-1] 目录
func (d *Diskv) pathFor(pathKey *PathKey) string {
	return filepath.Join(d.BasePath, filepath.Join(pathKey.Path...))
}

// 给定一个 PathKey 返回它所表示的文件的完整路径
// 也就是 <basedir>/Path[0]/Path[1]/.../Path[n-1]/FileName 路径
func (d *Diskv) completeFilename(pathKey *PathKey) string {
	return filepath.Join(d.pathFor(pathKey), pathKey.FileName)
}

// 确保存储键值对的目录存在
func (d *Diskv) ensurePathWithLock(pathKey *PathKey) error {
	return os.MkdirAll(d.pathFor(pathKey), d.PathPerm)
}

func (d *Diskv) createKeyFileWithLock(pathKey *PathKey) (*os.File, error) {
	if d.TempDir != "" {
		// 创建临时目录和临时文件
		if err := os.MkdirAll(d.TempDir, d.PathPerm); err != nil {
			return nil, fmt.Errorf("temp mkdir: %s", err)
		}
		f, err := os.CreateTemp(d.TempDir, "")
		if err != nil {
			return nil, fmt.Errorf("temp file: %s", err)
		}
		// 赋予 d.FilePerm 定义的权限
		if err := os.Chmod(f.Name(), d.FilePerm); err != nil {
			f.Close()
			// 发生错误时必须手动删除临时文件
			os.Remove(f.Name())
			return nil, fmt.Errorf("chmod: %s", err)
		}
		return f, nil
	}
	mode := os.O_WRONLY | os.O_CREATE | os.O_TRUNC // 只允许写入
	f, err := os.OpenFile(d.completeFilename(pathKey), mode, d.FilePerm)
	if err != nil {
		return nil, fmt.Errorf("open file: %s", err)
	}
	return f, nil
}

// 通过包装一个 io.Writer 来实现 io.WriteCloser 接口
type nopWriteCloser struct {
	io.Writer
}

func (wc *nopWriteCloser) Write(p []byte) (int, error) { return wc.Writer.Write(p) }
func (wc *nopWriteCloser) Close() error                { return nil }

func (d *Diskv) writeStreamWithLock(pathKey *PathKey, r io.Reader, sync bool) error {
	// 1. 创建目录
	if err := d.ensurePathWithLock(pathKey); err != nil {
		return fmt.Errorf("ensure path: %s", err)
	}
	// 2. 创建文件
	f, err := d.createKeyFileWithLock(pathKey)
	if err != nil {
		return fmt.Errorf("create key file: %s", err)
	}
	wc := io.WriteCloser(&nopWriteCloser{f})
	// 3. 是否进行压缩存储
	if d.Compression != nil {
		wc, err = d.Compression.Writer(f)
		if err != nil {
			f.Close()
			os.Remove(f.Name())
			return fmt.Errorf("compression writer: %s", err)
		}
	}
	// 4. 写入数据
	if _, err := io.Copy(wc, r); err != nil {
		f.Close()
		os.Remove(f.Name())
		return fmt.Errorf("i/o copy: %s", err)
	}
	if err := wc.Close(); err != nil {
		f.Close()
		os.Remove(f.Name())
		return fmt.Errorf("compression close: %s", err)
	}
	// 5. 是否将当前文件中的内容持久化到磁盘
	if sync {
		if err := f.Sync(); err != nil {
			f.Close()
			os.Remove(f.Name())
			return fmt.Errorf("file sync: %s", err)
		}
	}
	// 6. 关闭文件
	if err := f.Close(); err != nil {
		return fmt.Errorf("file close: %s", err)
	}
	fullPath := d.completeFilename(pathKey)
	if f.Name() != fullPath {
		if err := os.Rename(f.Name(), fullPath); err != nil {
			os.Remove(f.Name())
			return fmt.Errorf("rename: %s", err)
		}
	}
	// 7. 是否创建索引
	if d.Index != nil {
		d.Index.Insert(pathKey.originalKey)
	}
	// 8. 如果新写入的内容对应的键在缓存中存在，需要从缓存中删除这个键
	d.bustCacheWithLock(pathKey.originalKey)
	return nil
}

// bust: 打破，摔碎
func (d *Diskv) bustCacheWithLock(key string) {
	if val, ok := d.cache[key]; ok {
		d.uncacheWithLock(key, uint64(len(val)))
	}
}

func (d *Diskv) uncacheWithLock(key string, sz uint64) {
	d.cacheSize -= sz
	delete(d.cache, key)
}

// 将文件 srcFilename 中的内容设置为键 dstKey 对应的值
// 如果 dstKey 不存在则创建之；存在则覆盖其原来的值
// 如果 move 被设置为 true 则会删除 srcFilename 文件
func (d *Diskv) Import(srcFilename, dstKey string, move bool) (err error) {
	if dstKey == "" {
		return errEmptyKey
	}
	if fi, err := os.Stat(srcFilename); err != nil {
		return err
	} else if fi.IsDir() {
		return errImportDirectory
	}
	dstPathKey := d.transform(dstKey)
	d.mu.Lock()
	defer d.mu.Unlock()
	if err := d.ensurePathWithLock(dstPathKey); err != nil {
		return fmt.Errorf("ensure path: %s", err)
	}
	if move {
		if err := syscall.Rename(srcFilename, d.completeFilename(dstPathKey)); err == nil {
			d.bustCacheWithLock(dstPathKey.originalKey)
			return nil
		} else if err != syscall.EXDEV {
			return err
		}
	}
	f, err := os.Open(srcFilename)
	if err != nil {
		return err
	}
	defer f.Close()
	err = d.writeStreamWithLock(dstPathKey, f, false)
	if err == nil && move {
		err = os.Remove(srcFilename)
	}
	return err
}

func (d *Diskv) ReadString(key string) string {
	value, _ := d.Read(key)
	return string(value)
}

// 读取键 key 对应的值，发生错误时会返回空字符串
// 如果缓存命中则不会读取磁盘，否则会将键和对应的值存储到缓存中
func (d *Diskv) Read(key string) ([]byte, error) {
	rc, err := d.ReadStream(key, false)
	if err != nil {
		return []byte{}, err
	}
	defer rc.Close()
	return io.ReadAll(rc)
}

// 读取键 key 对应的值，如果缓存命中且 direct 为 false 则使用缓存中的值；否则从磁盘中读取
// 如果 direct 为 true 则会从缓存中删除键 key 和其对应的值
func (d *Diskv) ReadStream(key string, direct bool) (io.ReadCloser, error) {
	pathKey := d.transform(key)
	d.mu.RLock()
	defer d.mu.RUnlock()
	if val, ok := d.cache[key]; ok {
		if !direct {
			buf := bytes.NewReader(val)
			if d.Compression != nil {
				return d.Compression.Reader(buf)
			}
			return io.NopCloser(buf), nil
		}
		go func() { // 懒惰删除
			d.mu.Lock()
			defer d.mu.Unlock()
			d.uncacheWithLock(key, uint64(len(val)))
		}()
	}
	return d.readWithRLock(pathKey)
}

// 从磁盘中读取
func (d *Diskv) readWithRLock(pathKey *PathKey) (io.ReadCloser, error) {
	filename := d.completeFilename(pathKey)
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	if fi.IsDir() {
		return nil, os.ErrNotExist
	}
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var r io.Reader
	if d.CacheSizeMax > 0 {
		r = newSiphon(f, d, pathKey.originalKey)
	} else {
		r = &closingReader{f}
	}
	var rc = io.ReadCloser(io.NopCloser(r))
	if d.Compression != nil {
		rc, err = d.Compression.Reader(r)
		if err != nil {
			return nil, err
		}
	}
	return rc, nil
}

// newSiphon constructs a siphoning reader that represents the passed file.
// When a successful series of reads ends in an EOF, the siphon will write
// the buffered data to Diskv's cache under the given key.
func newSiphon(f *os.File, d *Diskv, key string) io.Reader {
	return &siphon{
		f:   f,
		d:   d,
		key: key,
		buf: &bytes.Buffer{},
	}
}

// siphon is like a TeeReader: it copies all data read through it to an
// internal buffer, and moves that buffer to the cache at EOF.
type siphon struct {
	f   *os.File
	d   *Diskv
	key string
	buf *bytes.Buffer
}

func (s *siphon) Read(p []byte) (int, error) {
	n, err := s.f.Read(p)
	if err == nil {
		return s.buf.Write(p[0:n])
	}
	if err == io.EOF {
		s.d.cacheWithoutLock(s.key, s.buf.Bytes())
		if closeErr := s.f.Close(); closeErr != nil {
			return n, closeErr
		}
		return n, err
	}
	return n, err
}

func (d *Diskv) cacheWithoutLock(key string, val []byte) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.cacheWithLock(key, val)
}

// cacheWithLock attempts to cache the given key-value pair in the store's
// cache. It can fail if the value is larger than the cache's maximum size.
func (d *Diskv) cacheWithLock(key string, val []byte) error {
	// If the key already exists, delete it.
	d.bustCacheWithLock(key)
	valueSize := uint64(len(val))
	if err := d.ensureCacheSpaceWithLock(valueSize); err != nil {
		return fmt.Errorf("%s; not caching", err)
	}
	// be very strict about memory guarantees
	if (d.cacheSize + valueSize) > d.CacheSizeMax {
		panic(fmt.Sprintf("failed to make room for value (%d/%d)", valueSize, d.CacheSizeMax))
	}
	d.cache[key] = val
	d.cacheSize += valueSize
	return nil
}

// ensureCacheSpaceWithLock deletes entries from the cache in arbitrary order
// until the cache has at least valueSize bytes available.
func (d *Diskv) ensureCacheSpaceWithLock(valueSize uint64) error {
	if valueSize > d.CacheSizeMax {
		return fmt.Errorf("value size (%d bytes) too large for cache (%d bytes)", valueSize, d.CacheSizeMax)
	}
	safe := func() bool { return (d.cacheSize + valueSize) <= d.CacheSizeMax }
	for key, val := range d.cache {
		if safe() {
			break
		}
		d.uncacheWithLock(key, uint64(len(val)))
	}
	if !safe() {
		panic(fmt.Sprintf("%d bytes still won't fit in the cache! (max %d bytes)", valueSize, d.CacheSizeMax))
	}
	return nil
}

// closingReader provides a Reader that automatically closes the
// embedded ReadCloser when it reaches EOF
type closingReader struct {
	rc io.ReadCloser
}

func (cr closingReader) Read(p []byte) (int, error) {
	n, err := cr.rc.Read(p)
	if err == io.EOF {
		if closeErr := cr.rc.Close(); closeErr != nil {
			return n, closeErr // close must succeed for Read to succeed
		}
	}
	return n, err
}
```

