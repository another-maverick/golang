package main

import (
	"net/http"

	prbuf "github.com/another-maverick/golang/protocolBuffer"
	proto "github.com/golang/protobuf/proto"
)

func main() {
	http.HandleFunc("/", prBufHandler)
	http.ListenAndServe(":12345", nil)
	/*
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

		fmt.Println(uRes.Name)
		fmt.Println(uRes.Email)
		fmt.Println(uRes.Id)
	*/
}

func prBufHandler(res http.ResponseWriter, req *http.Request) {
	u1 := &prbuf.User{
		Name:  proto.String("Stephen Curry"),
		Id:    proto.Int32(30),
		Email: proto.String("curry@warriors.com"),
	}
	body, err := proto.Marshal(u1)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type:", "application/x-protobuf")
	res.Write(body)
}
