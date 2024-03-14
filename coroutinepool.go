package main

import (
	"sync"
)

// 协程池

type Job struct {
	id    int
	input any
}

type Result struct {
	job   Job
	ouput any
}

const capacity = 5

// 全局任务队列
var jobs = make(chan Job, capacity)

// 向任务队列中添加任务
func Allocate(njobs int) {
	for i := 0; i < njobs; i++ {
		jobs <- Job{id: i, input: i}
	}
	close(jobs)
}

// 全局结果队列
var results = make(chan Result, capacity)

// 从结果队列中读取结果
func GetResult(done chan<- bool) {
	for result := range results {
		_, _, _ = result.job.id, result.job.input.(int), result.ouput.(int)
	}
	done <- true
}

// 启动 n 个协程并等待它们结束
func CreatePool(nworkers int) {
	var wg sync.WaitGroup
	for i := 0; i < nworkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

// 从任务队列中取出任务处理后发送到结果队列
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		results <- Result{job: job, ouput: work(job.input)}
	}
	wg.Done()
}

// 对输入进行处理产生输出的函数
func work(input any) any {
	return input
}
