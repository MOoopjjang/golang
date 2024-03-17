package main

import (
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age  int
	Addr string
}

var caches = []Person{
	Person{"xferlog", 10, "incheon"},
	Person{"kknda", 20, "seoul"},
	Person{"kcwda", 30, "yoong"},
}

func main() {

	// 1. 핸들러 등록
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
	mux.HandleFunc("/p", func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		name := values.Get("name")
		idx := -1

		for i, v := range caches {
			if v.Name == name {
				idx = i
				break
			}
		}

		if idx != -1 {
			fmt.Fprintf(w, "%s , %d , %s", caches[idx].Name, caches[idx].Age, caches[idx].Addr)
		} else {
			fmt.Fprintf(w, "%s not found", name)
		}
	})

	// 2. server start
	http.ListenAndServe(":8080", mux)

}
