package main

import "fmt"

type opFunc func(int, int) int

func calculate(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "-" {
		return func(a, b int) int {
			return a - b
		}
	} else {
		return nil
	}
}

func tst1() {

	for {
		var sel string
		fmt.Scan(&sel)

		exec := calculate(sel)
		result := exec(10, 20)
		fmt.Printf("result => %d \n", result)

	}

}

func main() {
	tst1()
}
