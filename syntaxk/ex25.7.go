package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			// Make a body
			car := &Car{}
			(*car).Body = "Sports car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func MakeTire(tireCh, paintCh chan *Car) {
	for c := range tireCh {
		// Make a Body
		time.Sleep(1 * time.Second)
		c.Tire = "Winter tire"
		paintCh <- c
	}
	wg.Done()
	close(paintCh)
}

func MakeColor(paintCh chan *Car) {
	for c := range paintCh {
		time.Sleep(1 * time.Second)
		c.Color = "red"
		//경과시간 출력
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car : %s %s %s \n", duration.Seconds(), c.Body, c.Tire, c.Color)

	}
	wg.Done()
}

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Println("Start Factory")

	wg.Add(3)
	go MakeBody(tireCh)
	go MakeTire(tireCh, paintCh)
	go MakeColor(paintCh)

	wg.Wait()
	fmt.Println("End Factory")
}
