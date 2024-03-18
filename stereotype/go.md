[toc]

# 基础

- 特点：简单、并发、垃圾回收

- 类型 `string` 和 `[]byte` 对比：不可变 `VS` 可变、Unicode 字符 `VS` 字节、可相互转换

- 数组 `VS` 切片：固定长度 `VS` 可变长度、值类型 `VS` 引用类型

- 切片扩容：小于 1024 时每次翻倍，大于 1024 时每次增长 1.25 倍

## 哈希表

- 哈希冲突的解决办法：开放寻址法、拉链法。对开放寻址法影响最大的因素是**负载因子**。

- Go 语言哈希表的特点

  - 使用拉链法解决哈希冲突

  - 桶的数量都是都是 2 的幂 -> `len(buckets) = 2 ^ hmap.B`
  - 通过 `hmap.hash0` 为哈希的结果引入随机性
  - 使用正常桶和溢出桶设计
  - 每个桶只能存储 8 个键值对，当前哈希的某个桶超出 8 个则把新的键值对存储到溢出桶中
  - 实现原理：散列表 -> 哈希函数 -> 存储桶 -> 链表 or 红黑树
  - 元素数量达到负载因子时进行扩容：容量翻倍 -> 分配 -> 重哈希 -> 切换新 -> 释放旧

## 函数

- 内置的 `panic` 和 `recover` 函数

  - 函数 `panic` 用于引发运行时错误，立即停止当前 goroutine 运行并按调用栈

    向上搜寻，执行每个调用者的 `defer` 函数，最终程序终止

  - 函数 `recover` 用于从 `panic` 引发的错误中恢复，捕获 `panic` 的值

- `defer` 函数：按 `LIFO` 顺序执行、函数参数预先计算、可以修改（有名返回值）函数的返回值

- 执行顺序：`import` -> `const` -> `var` -> `init()` -> `main()`

- 局部变量的逃逸分析：如果变量离开作用域后没有被引用，分配到栈上，否则分配到堆上

- 两个接口的比较：只有都为 `nil` 或类型 `T` 和值 `V` 同时相等时才相等

## 反射

核心类型是 `reflect.Type` 和 `reflect.Value` ，通过 `reflect.TypeOf` 和 `reflect.ValueOf` 将

**任意类型的变量**转换成这两种类型然后进行各种操作。

通过反射调用方法：

```go
func main() {
	a := animal{}
	v := reflect.ValueOf(a)
	v.MethodByName("Run").Call(make([]reflect.Value, 0))
}

type animal struct{}

func (a animal) Run() {
	fmt.Println("Run")
}
```

# 并发

## 锁

- 协程是轻量级的用户态线程，由 Go 调度器进行管理。协程之间通过管道通信

- `Mutex` 的两种模式：正常（FIFO）和饥饿（一定的公平原则）模式

  **正常模式**：加锁时如果当前 `Locked` 位为 1 ，说明该锁当前由其他协程持有，尝试加锁的协程并不是马上转入

  阻塞，而是会**持续的探测 `Locked` 位是否变为 0 ，这个过程即为自旋过程。自旋时间很短，但如果在自旋过程中**

  **发现锁已被释放，那么协程可以立即获取锁。此时即便有协程被唤醒也无法获取锁，只能再次阻塞**。自旋的

  好处是，当加锁失败时不必立即转入阻塞，有一定机会获取到锁，这样可以避免协程的切换。

  **饥饿模式**：被唤醒的协程得到 CPU 后开始运行，此时发现锁已被抢占了，自己只好再次阻塞，不过阻塞前会

  判断自上次阻塞到本次阻塞经过了多长时间，**如果超过 1ms 的话，会将 `Mutex` 标记为饥饿模式，然后再阻塞**。

  **处于饥饿模式下，不会启动自旋过程**，也即一旦有协程释放了锁，那么一定会唤醒协程，被唤醒的

  协程将会成功获取锁，同时也会把等待计数减 1 。

## 上下文

- 上下文 `context.Context` 的最大作用：<font color=red>**在 goroutine 构成的树形结构中对信号进行同步以减少计算资源的浪费**</font>。

- 上下文 `context.Context` 的使用方法：<font color=red>**多个 goroutine 同时订阅 `ctx.Done()` 管道中的消息，一旦接收到**</font>

  <font color=red>**取消信号就立刻停止当前正在执行的工作**</font>。

  ```go
  func main() {
      // 过期时间为一秒的上下文
  	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
  	defer cancel()
  	// 处理时间 500 毫秒触发 <-time.After(duration) 分支
      // 输出
      // process request with 500ms 处理线程完成
  	// main context deadline exceeded 主线程超时
  	go handle(ctx, 500*time.Millisecond)
  	<-ctx.Done()
  	fmt.Println("main", ctx.Err())
  }
  
  func main() {
      // 过期时间为一秒的上下文
  	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
  	defer cancel()
  	// 处理时间 1500 毫秒触发 <-ctx.Done() 分支
      // 输出
      // main context deadline exceeded 主线程超时
  	// handle context deadline exceeded 处理线程超时
  	go handle(ctx, 1500*time.Millisecond)
  	<-ctx.Done()
  	fmt.Println("main", ctx.Err())
  }
  
  func handle(ctx context.Context, duration time.Duration) {
  	select {
  	case <-time.After(duration):
  		fmt.Println("process request with", duration)
  	case <-ctx.Done():
  		fmt.Println("handle", ctx.Err())
  	}
  }
  ```

# 垃圾收集

Java 的垃圾回收：线性、并发、并行标记清除、G1

Go 的垃圾回收：标记清除法 -> 1.3 -> 三色标记法 -> 1.8 -> 三色标记法 + 混合写屏障

# 调度

总体描述：

**m 个 G 要分配到 n 个 M 上运行，M 最多有 `GOMAXPROCS` 个，每个 M 依附于一个 P ，每个 P 在同一时刻只能运行一个 M ，如果 P 上的 M 阻塞就需要其他 M 来运行**。

调度器的核心目标：**将 goroutine 调度到内核线程上运行**。

调度器的核心思想：**(G)重用线程(M)**、**限制同时运行 `GOMAXPROCS` 个线程(M)**、**(P)线程私有运行队列(存放G)**。

Go 语言的调度器通过使用与 CPU 数量相等的线程减少线程频繁切换的内存开销，同时在每一个线程上执行额外开销

更低的 goroutine 来降低操作系统和硬件的负载。

数据结构：

- G — 表示 goroutine，它是一个待执行的**任务**。它的状态可以被划分为 3 类：**等待中、可运行、运行中**。
- M — 表示**操作系统的线程**，它由操作系统的调度器调度和管理。
- P — 表示处理器，它可以被看做**运行在线程上的本地调度器**。通过调度**每一个内核线程都能够执行多个 goroutine**。

运行队列：

每个 P 都维护一个由可运行的 G 组成的队列。线程阻塞后可以将 P 上的 G 转移到其他线程。

除了**本地可运行队列 LRQ**外，还有一个**全局可运行队列GRQ**，GRQ 中存储的 G 不对应具体的 P 。

调度时机：

关键字 `go` 、GC、系统调用、内存同步访问

调度器的工作：

每一轮找到可运行的 G 并执行，顺序：LRQ -> GRQ -> net poll -> 从其他 P 偷取
