package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":"+os.Getenv("TESTPORT"), nil)

}
func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Println(res, "This is the Homepage.")
}
