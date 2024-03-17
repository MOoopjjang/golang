package main

import (
	"fmt"
	"os"
)

func main() {
	size := len(os.Args)
	fmt.Printf("size = %d \n", size)

	for i, v := range os.Args {
		fmt.Printf("[%d]%s \n", i, v)
	}
}
