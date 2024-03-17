package main

import (
	"fmt"
	"sync"
	"time"
)

func squre(wg *sync.WaitGroup, ch chan int) {
	// 채널에 데이타가 들어올때까지 wait
	c := <-ch
	time.Sleep(1 * time.Second)
	fmt.Printf(" c = %d \n", c)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	var ch chan int = make(chan int)
	wg.Add(1)
	go squre(&wg, ch)
	ch <- 9
	wg.Wait()

}
