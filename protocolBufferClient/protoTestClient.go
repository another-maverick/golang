package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	prbuf "github.com/another-maverick/golang/protocolBuffer"
	proto "github.com/golang/protobuf/proto"
)

func main() {

	myres, err := http.Get("http://localhost:12345")
	if err != nil {
		fmt.Println("Cannot get the resource on port 12345")
	}
	defer myres.Body.Close()

	b, err := ioutil.ReadAll(myres.Body)
	if err != nil {
		fmt.Println("cant read the response from localhost")
	}

	var uRes prbuf.User

	err = proto.Unmarshal(b, &uRes)
	if err != nil {
		fmt.Println("cant unmarshal the response body")
	}

	fmt.Println(uRes.GetName())
	fmt.Println(uRes.GetEmail())
	fmt.Println(uRes.GetId())
}
