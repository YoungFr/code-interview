package main

import (
	"fmt"
)

func main() {
	// 输出
	// 4
	// 3
	// 2
	// 1
	// 0
	for i := range 5 {
		defer func() {
			fmt.Println(i)
		}()
	}

	// 两个线程打印 1A2B3C4D...
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
}
