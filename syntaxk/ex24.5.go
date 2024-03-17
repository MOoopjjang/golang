package main

// dead lock example

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func dinigProblem(name string, first, second *sync.Mutex, firstname, secondname string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다. \n", name)
		first.Lock()
		fmt.Printf("%s %s 획득 \n", name, firstname)
		second.Lock()
		fmt.Printf("%s %s 획득 \n", name, secondname)

		fmt.Printf("%s 밥을 먹습니다 \n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go dinigProblem("A", fork, spoon, "포크", "수저")
	go dinigProblem("B", spoon, fork, "수저", "포크")
	wg.Wait()
}
