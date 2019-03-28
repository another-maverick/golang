package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye/", goodbye)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":12345", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Stephen Curry"
	}
	fmt.Fprint(res, "Hello , I am ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Stephen Curry"
	}
	fmt.Fprint(res, "Goodbye from ", name)
}
func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "This is the homepage!")
}
