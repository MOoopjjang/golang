package main

import (
	"fmt"
	"os"
)

// func vargs(param ...interface{}) {
// 	for _, arg := range param {
// 		switch ff := arg.(type) {
// 		case string:
// 			vv := ff.(string)
// 			fmt.Println(vv)
// 		case int:
// 			vv := ff.(int)
// 			fmt.Println(vv)
// 		case float64:
// 			vv := ff.(int)
// 			fmt.Println(vv)
// 		}
// 	}
// }

func createFile(fileName string) {
	fmt.Println("fileName : ", fileName)

	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("file create failed")
		return
	}

	defer fmt.Println("반드시 호출됩니다")
	defer f.Close()
	defer fmt.Println("파일을 닫습니다")

	fmt.Println("파일에 Hello World를 씁니다.")
	fmt.Fprintln(f, "Hello World")
}

func main() {
	// vargs("xferlog", 10, 20.1444)

	createFile("hello.txt")

}
