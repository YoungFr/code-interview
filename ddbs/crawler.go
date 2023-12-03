package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns a slice of URLs found on the page.
	Fetch(url string) (urls []string, err error)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) ([]string, error) {
	if res, ok := f[url]; ok {
		fmt.Printf("found:   %s\n", url)
		return res.urls, nil
	}
	fmt.Printf("missing: %s\n", url)
	return nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

// Serial crawler

func Serial(url string, fetcher Fetcher, fetched map[string]bool) {
	if fetched[url] {
		return
	}
	fetched[url] = true
	urls, err := fetcher.Fetch(url)
	if err != nil {
		return
	}
	for _, u := range urls {
		Serial(u, fetcher, fetched)
	}
}

// Concurrent crawler with shared state and Mutex

// 所有 goroutine 共享同一个 fetched 哈希表
// 使用锁来保护它
type fetchState struct {
	mu      sync.Mutex
	fetched map[string]bool
}

func makeState() *fetchState {
	return &fetchState{fetched: make(map[string]bool)}
}

// 如果线程 T1 和 T2 使用同一个 url 读取 fetched[url] 的值
// 假设这个 url 还没有被访问过 => 线程 T1 和 T2 都读到 false 值
// 两个线程就会对同一个 url 进行 Fetch 调用
func (fs *fetchState) testAndSet(url string) bool {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	r := fs.fetched[url]
	fs.fetched[url] = true
	return r
}

func ConcurrentMutex(url string, fetcher Fetcher, fs *fetchState) {
	if fs.testAndSet(url) {
		return
	}
	urls, err := fetcher.Fetch(url)
	if err != nil {
		return
	}
	// 使用 sync.WaitGroup 等待所有 goroutine 结束
	var done sync.WaitGroup
	for _, url := range urls {
		done.Add(1)
		u := url
		go func(u string) {
			defer done.Done()
			ConcurrentMutex(u, fetcher, fs)
		}(u)
	}
	done.Wait()
}

// Concurrent crawler with channels

func worker(url string, ch chan []string, fetcher Fetcher) {
	urls, err := fetcher.Fetch(url)
	if err != nil {
		ch <- make([]string, 0)
	} else {
		ch <- urls
	}
}

func coordinator(ch chan []string, fetcher Fetcher) {
	n := 1
	fetched := make(map[string]bool)
	// 如果接收操作先执行
	// 接收方 goroutine 会阻塞直到另一个 goroutine 发送一个值
	for urls := range ch {
		for _, u := range urls {
			if !fetched[u] {
				fetched[u] = true
				n += 1
				go worker(u, ch, fetcher)
			}
		}
		n -= 1
		if n == 0 {
			break
		}
	}
}

func ConcurrentChannel(url string, fetcher Fetcher) {
	// 无缓冲管道
	ch := make(chan []string)
	go func() {
		// 无缓冲管道的发送操作会阻塞
		// 直到另一个 goroutine 在对应的通道上执行接收操作
		ch <- append([]string(nil), url)
	}()
	coordinator(ch, fetcher)
}

func main() {
	fmt.Printf("=== Serial===\n")
	Serial("http://golang.org/", fetcher, make(map[string]bool))

	fmt.Printf("=== ConcurrentMutex ===\n")
	ConcurrentMutex("http://golang.org/", fetcher, makeState())

	fmt.Printf("=== ConcurrentChannel ===\n")
	ConcurrentChannel("http://golang.org/", fetcher)
}
