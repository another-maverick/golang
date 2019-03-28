package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/fooServer", fooServer)
	http.ListenAndServe("localhost:1234", nil)
}

func fooServer(res http.ResponseWriter, req *http.Request) {

	fmt.Fprint(res, "My name is Jon Snow")
}
