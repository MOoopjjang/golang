package main

import (
	"fmt"
	"strings"
)

var dictArr [3]string = [3]string{"hahaha king my life", "kimcw is programmer", "73892 akdls xldflk!!!"}

var intArr [7]int = [7]int{1, 4, 2, 6, 6, 9, 10}

func search(text string) int {
	for i, v := range dictArr {
		if strings.Contains(v, text) {
			return i
		}
	}

	return -1
}

func searchInt(num int) (int, error) {
	fmt.Printf("search num = %d \n", num)

	for i, v := range intArr {
		if v == num {
			return i, nil
		}
	}

	return -1, fmt.Errorf("존재하지 않는 값입니다")
}

func main() {
	sel := 0
	fmt.Println("***************************************")
	fmt.Println("1. 숫자 search ")
	fmt.Println("2. 문자 검색")
	fmt.Scan(&sel)
	fmt.Println("***************************************")
	fmt.Println(" 값을 입력하세요::")
	if sel == 1 {
		num := 0
		fmt.Scan(&num)
		idx, err := searchInt(num)
		if err == nil {
			fmt.Printf(">>> index = %d \n", idx)
		} else {
			fmt.Println(err)
		}
	} else {
		text := ""
		fmt.Scan(&text)
		idx := search(text)
		fmt.Printf(">>> text -> idx = %d \n", idx)
	}
}
