package main

//====================================================
// 지정한 위치의 디렉토리및 하위디렉토리에서 .txt .log 파일을 읽어서
//  파일 내에 특정단어가 나는 위치를 찾는 프로그램
//----------------------------------------------------
// step 1 ( 완료 )
//  - 디렉토리 순회기능
//  - file open / read기능
//  - go-routine / channel
//======================================================
// step 2
//   - txt파일을 목록을 저장
//   - 저장된 txt파일 목록만큼 go-routine 생성
//
//======================================================

import (
	"bufio"
	"fmt"

	// "goproject/demo/word/search"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type SearchInfo struct {
	RootDir  string
	FindWord string
}

func (si *SearchInfo) ToStringn() string {
	return fmt.Sprintf("start dir=%s , find = %s", (*si).RootDir, (*si).FindWord)
}

var arTxt [2]string = [...]string{"txt", "log"}
var wg sync.WaitGroup

func isDir(dirPath string) bool {
	if stat, err := os.Stat(dirPath); err == nil {
		return stat.IsDir()
	} else {
		return false
	}
}

func isTxt(f string) bool {
	r := []rune(f)
	idx := -1
	for i := len(r) - 1; i >= 0; i-- {
		if r[i] == '.' {
			idx = i
			break
		}
	}

	if idx != -1 {
		if rr := r[idx+1:]; len(rr) == 3 {
			for _, vv := range arTxt {
				if vv == string(rr) {
					return true
				}
			}

		}

	}
	return false
}

func recursive(recursiveCh chan *SearchInfo, readCh chan *SearchInfo) {

	for ch := range recursiveCh {
		d := ch.RootDir
		keyword := ch.FindWord
		err := filepath.Walk(d, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return err
			}

			if !info.IsDir() && isTxt(path) {
				fmt.Printf("%s \n", path)
				// ReadFile(keyword, path)
				nsearchInfo := &SearchInfo{path, keyword}
				readCh <- nsearchInfo

			}
			return nil
		})

		if err != nil {
			panic(err)
		}
	}
	close(readCh)
	wg.Done()

}

func GetFindDirPath() (bool, SearchInfo) {
	searchInfo := SearchInfo{}
	var inputPath string
	fmt.Println("검색할 위치를 입력하세요 ( 종료할려면 q ):")
	fmt.Scan(&inputPath)
	searchInfo.RootDir = inputPath
	if inputPath == "q" {
		return false, searchInfo
	} else {
		if !isDir(inputPath) {
			return false, searchInfo
		}

		fmt.Println("검색할 단어는:")
		fmt.Scan(&inputPath)
		searchInfo.FindWord = inputPath
	}
	return true, searchInfo
}

func ReadFile(readCh chan *SearchInfo) {
	for ch := range readCh {
		if fo, err := os.Open(ch.RootDir); err != nil {
			defer fo.Close()
			panic(err)
		} else {
			defer fo.Close()
			reader := bufio.NewReader(fo)
			fmt.Printf("---------- %s ------------\n", ch.RootDir)
			idx := 1
			for {
				line, prefix, err := reader.ReadLine()
				if prefix || err != nil {
					break
				}

				if IsFindWord(ch.FindWord, string(line)) {
					fmt.Printf("[%d] %s \n", idx, string(line))
				}
				idx++

			}
		}
	}
	wg.Done()
}

func IsFindWord(keyWord string, line string) bool {
	return strings.Contains(line, keyWord)
}

func main() {
	wg.Add(2)
	recursiveCh := make(chan *SearchInfo)
	readCh := make(chan *SearchInfo)

	go recursive(recursiveCh, readCh)
	go ReadFile(readCh)

	for {
		if ret, searchInfo := GetFindDirPath(); !ret {
			fmt.Printf("%s invalid \n", searchInfo.RootDir)
			break
		} else {
			recursiveCh <- &searchInfo
		}
	}

	close(recursiveCh)
	wg.Wait()

}
