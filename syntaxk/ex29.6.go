package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	student := Student{"cwkim", 10, 10}
	// Student 객체를 []byte로 변환
	data, _ := json.Marshal(student)
	// JSON포멧임을 표시
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// 결과 전송
	fmt.Fprintf(w, string(data))
}

func main() {
	http.ListenAndServe(":8080", MakeHandler())
}
