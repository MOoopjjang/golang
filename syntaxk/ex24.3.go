package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Accounnt struct {
	Balance int
}

func DepositAndWithdraw(account *Accounnt) {
	mutex.Lock()
	defer mutex.Unlock()
	if (*account).Balance < 0 {
		fmt.Println("not calculate balanced :", (*account).Balance)
	}

	(*account).Balance += 1000
	time.Sleep(time.Millisecond)
	(*account).Balance -= 1000

	fmt.Printf(" balance = %d \n", (*account).Balance)
}

func main() {
	var wg sync.WaitGroup
	account := &Accounnt{0}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			for {

				DepositAndWithdraw(account)

			}
			wg.Done()
		}()
	}

	wg.Wait()
}
