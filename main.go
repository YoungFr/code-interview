package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1. defer 的执行顺序

	// Output: 43210
	defer func() {
		fmt.Println()
	}()
	for i := range 5 {
		defer func() {
			fmt.Print(i)
		}()
	}

	// 2. 两个线程打印 1A2B3C4D...

	var (
		number = make(chan int)
		done   = make(chan int)
	)
	go func() {
		for i := 1; i <= 26; i++ {
			number <- i
		}
	}()
	go func() {
		for c := 'A'; c <= 'Z'; c++ {
			fmt.Print(<-number, " ", string(c), " ")
			if c == 'Z' {
				done <- 1
			}
		}
	}()
	<-done
	fmt.Println()

	// 3. 启动 N 个线程按顺序打印 0 ~ N-1 数字

	N := 30

	// 3.1 锁
	var count int     // 多个线程争抢 count 变量和锁
	var mu sync.Mutex // 多个线程争抢 count 变量和锁
	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		wg.Add(1)
		// 每个线程记录自己要打印的值
		go func(i int) {
			// 不断地争抢锁
			for {
				mu.Lock()
				now := count
				// 如果争抢到锁并且当前 count 的值与自己要打印的相同则打印
				// 打印完成后将 count 加一再释放锁后退出
				if now == i {
					fmt.Print(now, " ")
					count++
					mu.Unlock()
					wg.Done()
					return
				}
				// 争抢到了锁但是当前 count 的值与自己要打印的不同
				// 释放锁然后继续争抢
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()
	fmt.Println()

	// 3.2 管道
	quit := make(chan int)
	chs := make([]chan int, N)
	for i := 0; i < N; i++ {
		chs[i] = make(chan int)
	}
	for i := 0; i < N; i++ {
		go func(i int) {
			fmt.Print(<-chs[i], " ")
			// 打印完最后一个数字后通知主线程
			if i == N-1 {
				quit <- 1
				return
			}
			chs[i+1] <- i + 1
		}(i)
	}
	go func() {
		chs[0] <- 0
	}()
	<-quit // 主线程等待
	fmt.Println()

	// 4. 按照 零 -> 奇数 -> 偶数 -> 零 -> 奇数 -> 偶数 的顺序打印数字

	var (
		zero     = make(chan struct{})
		odd      = make(chan struct{})
		even     = make(chan struct{})
		taskdone = make(chan struct{})
	)
	go printZero(10, zero, odd)
	go printOdd(10, odd, even)
	go printEven(10, even, zero, taskdone)
	zero <- struct{}{}
	<-taskdone
	fmt.Println()

	// 协程池的使用
	start := time.Now().UnixMicro()

	njobs := 10000
	go Allocate(njobs)

	d := make(chan bool)
	go GetResult(d)

	nworkers := 4
	CreatePool(nworkers)

	<-d

	end := time.Now().UnixMicro()
	fmt.Printf("%dus\n", end-start)

	// 1. 脏写
	// a := 0
	// var wg sync.WaitGroup
	// for i := 0; i < 1000; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		a += 10
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println(a)

	// 2. 读未提交：写之前加写锁 -> 脏读
	// a := 300
	// b := 700
	// var mu sync.RWMutex
	// var wg sync.WaitGroup
	// wg.Add(2)

	// // 写事务
	// go func() {
	// 	mu.Lock()
	// 	// 修改
	// 	a += 200
	// 	b -= 200
	// 	// 回滚
	// 	b += 200
	// 	a -= 200
	// 	mu.Unlock()
	// 	// 提交
	// 	wg.Done()
	// }()

	// // 读事务
	// go func() {
	// 	fmt.Printf("read T: a = %d, b = %d\n", a, b)
	// 	// 提交
	// 	wg.Done()
	// }()

	// wg.Wait()
	// fmt.Printf("after Ts: a = %d, b = %d\n", a, b)

	// 3. 读已提交：写之前加写锁、读之前加读锁、读后释放读锁 -> 不可重复读
	// a := 300
	// b := 700
	// var mu sync.RWMutex
	// var wg sync.WaitGroup
	// wg.Add(2)

	// // 写事务
	// go func() {
	// 	mu.Lock()
	// 	a += 200
	// 	b -= 200
	// 	mu.Unlock()
	// 	// 提交
	// 	wg.Done()
	// }()

	// // 读事务
	// go func() {
	// 	mu.RLock()
	// 	fmt.Printf("1st read: a = %d, b = %d\n", a, b)
	// 	mu.RUnlock()
	// 	mu.RLock()
	// 	fmt.Printf("2nd read: a = %d, b = %d\n", a, b)
	// 	mu.RUnlock()
	// 	// 提交
	// 	wg.Done()
	// }()

	// wg.Wait()
	// fmt.Printf("after Ts: a = %d, b = %d\n", a, b)

	// 4. 可重复读：写之前加写锁、读之前加读锁 -> 幻读
	// a := 300
	// b := 700
	// var mu sync.RWMutex
	// var wg sync.WaitGroup
	// wg.Add(2)

	// // 写事务
	// go func() {
	// 	mu.Lock()
	// 	a += 200
	// 	b -= 200
	// 	mu.Unlock()
	// 	// 提交
	// 	wg.Done()
	// }()

	// // 读事务
	// go func() {
	// 	mu.RLock()
	// 	fmt.Printf("1st read: a = %d, b = %d\n", a, b)
	// 	fmt.Printf("2nd read: a = %d, b = %d\n", a, b)
	// 	mu.RUnlock()
	// 	// 提交
	// 	wg.Done()
	// }()

	// wg.Wait()
	// fmt.Printf("after Ts: a = %d, b = %d\n", a, b)

	// 单生产者多消费者
	tasks := make(chan int, 2)
	var pcwg sync.WaitGroup
	go produce(tasks, 5)
	// 启动 5 个消费者
	for i := 0; i < 5; i++ {
		pcwg.Add(1)
		go consume(i, tasks, &pcwg)
	}
	pcwg.Wait()
}

// 向队列 tasks 中生产 n 个任务
func produce(tasks chan<- int, n int) {
	for i := 0; i < n; i++ {
		tasks <- i
	}
	close(tasks)
}

// 编号为 id 的消费者从队列 tasks 中消费任务
func consume(id int, tasks <-chan int, wg *sync.WaitGroup) {
	for task := range tasks {
		fmt.Printf("consumer %d consumes task %d\n", id, task)
	}
	wg.Done()
}

func printZero(n int, zero <-chan struct{}, odd chan<- struct{}) {
	for i := 0; i < n; i++ {
		<-zero
		fmt.Print(0, " ")
		odd <- struct{}{}
	}
}

func printOdd(n int, odd <-chan struct{}, even chan<- struct{}) {
	for i := 1; i <= 2*n-1; i += 2 {
		<-odd
		fmt.Print(i, " ")
		even <- struct{}{}
	}
}

func printEven(n int, even <-chan struct{}, zero chan<- struct{}, done chan<- struct{}) {
	for i := 2; i <= 2*n; i += 2 {
		<-even
		fmt.Print(i, " ")
		if i == 2*n {
			done <- struct{}{}
			return
		}
		zero <- struct{}{}
	}
}
