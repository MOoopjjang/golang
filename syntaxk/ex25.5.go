package main

import (
	"fmt"
	"sync"
	"time"
)

func Square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("Square n = %d \n", n*n)
			time.Sleep(1 * time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int)
	quit := make(chan bool)
	go Square(&wg, ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i * 2

	}

	quit <- true
	wg.Wait()
}
