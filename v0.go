package main

import (
	"fmt"
	"net/http"
)

type rating struct {
	datas map[string]int
}

func (r *rating) counter(res http.ResponseWriter, req *http.Request) {
	pageName := req.URL.Query().Get("page")
	if _, ok := r.datas[pageName]; ok {
		r.datas[pageName]++
		fmt.Fprintf(res, "Done\n")
	}
}

func main() {
	r := rating{datas: map[string]int{"page_1": 0}}
	http.HandleFunc("/count", r.counter)
	http.ListenAndServe("localhost:8080", nil)
}
