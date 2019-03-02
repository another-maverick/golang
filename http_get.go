package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("http://golang.org")
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response body is ", string(body))
	response.Body.Close()

}
