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

// func (p *Person) ToString() {
// 	fmt.Fprintf("name = %s , age = %d , addr = %s", (*p).Name, (*p).Age, (*p).Addr)
// }

var caches = []Person{
	Person{"xferlog", 10, "incheon"},
	Person{"kknda", 20, "seoul"},
	Person{"kcwda", 30, "yoong"},
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	fmt.Printf("name = %s \n", name)

	idx := -1
	for i, v := range caches {
		if v.Name == name {
			idx = i
			break
		}
	}

	if idx != -1 {
		fmt.Fprintf(w, "name = %s , Age = %d , Addr = %s", caches[idx].Name, caches[idx].Age, caches[idx].Addr)
	}

}

func main() {
	http.HandleFunc("/person", personHandler)

	http.ListenAndServe(":8080", nil) //웹서버 시작
}
