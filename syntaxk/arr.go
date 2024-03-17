package main

import (
	"fmt"
)

func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10}}

	t := -1
	for r, one := range a {
		for _, v := range one {
			if t == -1 {
				t = r
			} else {
				if t != r {
					fmt.Println("##########################")
					t = r
				}
			}
			fmt.Printf("%d \n", v)
		}
	}
}
