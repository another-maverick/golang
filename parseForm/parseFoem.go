package main

import (
	"fmt"
	"net/http"
)

func parseMe(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:12345",
	}
	http.HandleFunc("/parseme", parseMe)
	server.ListenAndServe()
}
