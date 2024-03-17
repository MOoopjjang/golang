package main

import (
	"fmt"
	"sync"
)

// waitGroup 객체
var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}

	fmt.Printf("%d 에서 %d 까지 합은 %d 입니다.\n", a, b, sum)

	//작업이 완료됨을 표시
	wg.Done()
}

func main() {
	// 총 작업할 go-routine 개수 설정
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000000)
	}

	// 모든 작업이 완료되길 기다림
	wg.Wait()
}
