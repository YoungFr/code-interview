package main

import (
	"fmt"
	"sync"
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
			fmt.Print(<-number, string(c))
			if c == 'Z' {
				done <- 1
			}
		}
	}()
	<-done
	fmt.Println()

	// 3. 启动 N 个线程按顺序打印 0 ~ N-1 数字

	N := 50

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
					fmt.Print(now)
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
			fmt.Print(<-chs[i])
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
}
