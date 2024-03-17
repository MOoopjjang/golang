package main

import (
	"log"
	"os"
)

func init() {
	// 표준출력으로 로그를 출하도록 변경한다
	log.SetOutput(os.Stdout)
}

func main() {
	//지정된 검색어로 검색을 수행한다
	search.Runn("Sherlock Holmes")
}
