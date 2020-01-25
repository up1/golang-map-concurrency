package main

import (
	"fmt"
	"net/http"
)

type server struct {
	result chan<- rating
}

type rating struct {
	key       string
	value     int
	replyChan chan int
}

func manageData(i map[string]int) chan<- rating {
	datas := make(map[string]int)
	for k, v := range i {
		datas[k] = v
	}

	rs := make(chan rating)
	go func() {
		for r := range rs {
			datas[r.key]++
			r.replyChan <- datas[r.key]
		}
	}()
	return rs
}

func (s *server) counter(res http.ResponseWriter, req *http.Request) {
	pageName := req.URL.Query().Get("page")
	replyChan := make(chan int)
	s.result <- rating{key: pageName, replyChan: replyChan}
	reply := <-replyChan
	fmt.Fprintf(res, fmt.Sprintf("Done with %d\n", reply))
}

func main() {
	s := server{manageData(map[string]int{"page_1": 0})}
	http.HandleFunc("/count", s.counter)
	http.ListenAndServe("localhost:8080", nil)
}
