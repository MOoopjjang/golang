package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행인수가 필요합니다.")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	fmt.Println("찾을려는 단어:", word)
	PrintAllFiles(files)
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles(files []string) {
	for _, path := range files {
		filelist, err := GetFileList(path)
		if err != nil {
			fmt.Println("파일경로가 잘못되었습니다.")
			return
		}

		fmt.Println("현재 탐색할 파일리스트")
		for _, name := range filelist {
			fmt.Println(name)
		}
	}
}
