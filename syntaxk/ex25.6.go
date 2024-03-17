package main

import (
	"fmt"
	"sync"
	"time"
)

func Square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case n := <-ch:
			fmt.Printf("n = %d \n", n*n)
			time.Sleep(1 * time.Second)
		case <-tick:
			fmt.Printf("tick \n")
		case <-terminate:
			wg.Done()
			return
		}
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int)
	go Square(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}
