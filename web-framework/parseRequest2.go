package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request){
	h := r.Header["Accept-Language"]
	h1 := r.Header.Get("Accept-Language")
	fmt.Fprintln(w, h)
	fmt.Fprintln(w, h1)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:12345",
	}
	http.HandleFunc("/handle", headers)
	server.ListenAndServe()
}