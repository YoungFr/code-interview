package main

import (
	"math/rand"
	"sync"
)

// 协程池是一组随时等待分配任务的协程
// 每个协程接到任务就执行
// 执行完后继续等待下一次分配

type Job struct {
	id    int
	input any
}

type Result struct {
	job   Job
	ouput any
}

const capacity = 5

var jobs = make(chan Job, capacity)

func Allocate(njobs int) {
	for i := 0; i < njobs; i++ {
		jobs <- Job{id: i, input: rand.Int()}
	}
	close(jobs)
}

var results = make(chan Result, capacity)

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		results <- Result{job: job, ouput: work(job.input)}
	}
	wg.Done()
}

func work(input any) any {
	return input
}

func CreatePool(nworkers int) {
	var wg sync.WaitGroup
	for i := 0; i < nworkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func GetResult(done chan bool) {
	for result := range results {
		_, _, _ = result.job.id, result.job.input.(int), result.ouput.(int)
	}
	done <- true
}
