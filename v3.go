package main

import (
	"fmt"
	"net/http"
	"sync"
)

type rating struct {
	datas sync.Map
}

func (r *rating) counter(res http.ResponseWriter, req *http.Request) {
	pageName := req.URL.Query().Get("page")
	if value, ok := r.datas.Load(pageName); ok {
		r.datas.Store("page_1", int(value)+1)
		fmt.Fprintf(res, fmt.Sprintf("Done with %d\n", int(value)+1))
	}
}

func main() {
	r := rating{}
	r.datas.Store("page_1", 0)
	http.HandleFunc("/count", r.counter)
	http.ListenAndServe("localhost:8080", nil)
}
