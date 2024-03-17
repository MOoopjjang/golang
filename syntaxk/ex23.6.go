package main

import (
	"fmt"
)

func f() {
	fmt.Println("f() start")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}
	}()

	g()
	fmt.Println("f() end")
}

func g() {
	fmt.Printf("9 / 3  = %d \n", h(9, 3))
	fmt.Printf("9 / 0 = %d \n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("지수는 0일수 없습니다.")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램 계속 실행됨")
}
