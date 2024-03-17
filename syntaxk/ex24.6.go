package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SquareJob struct {
	Index int
}

func (job *SquareJob) Do() {
	fmt.Printf("%d 작업시작\n", job.Index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업완료 = %d\n", job.Index, job.Index*job.Index)
}

func main() {
	var jobs [10]Job

	for i := 0; i < 10; i++ {
		jobs[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobs[i]
		go func() {
			job.Do()
			wg.Done()
		}()

	}

	wg.Wait()

}
