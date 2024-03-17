package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	h := []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	for _, v := range h {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c", v)
	}
}

func PrintNum() {
	for i := 0; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func main() {
	go PrintHangul()
	go PrintNum()

	time.Sleep(10 * time.Second)
}
