package main

import "fmt"
import "net/http"

func MainHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(rw)
	fmt.Println(req.Body)
}

func TestHandler(rw http.ResponseWriter, req *http.Request) {
	// test
	rw.Write([]byte("test"))
}

func main() {
	RouteConfig := map[string]func(rw http.ResponseWriter, req *http.Request){
		"/":      MainHandler,
		"/test/": TestHandler,
	}
	for k, v := range RouteConfig {
		http.HandleFunc(k, v)
	}
	http.ListenAndServe(":9090", nil)
}
