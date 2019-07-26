package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author   string  `json:"author"`
 }


func main() {
	server := http.Server{
		Addr: "127.0.0.1:12345",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}


// handleRequest - gets the request and then invokes the right handler
func handleRequest(w http.ResponseWriter, r *http.Request){
	var err error
	switch r.Method{
	case "GET":
		err = handleGetRequest(w, r)
	case "POST":
		err = handlePostRequest(w,r)
	case "PUT":
		err = handlePutRequest(w,r)
	case "DELETE":
		err = handleDeleteRequest(w,r)
	}
	if err != nil {
		fmt.Println("Cannot handle the Method in the request")
		return
	}

}

// Get a record from the database
func handleGetRequest(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		fmt.Println("cannot parse the URL in the GET request")
		return
	}

	post, err := retrieve(id)
	if err != nil {
		fmt.Println("cannot retrieve the data from database")
		return
	}

	output, err := json.MarshalIndent(&post, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}


// Post request to ther webservice - normally stores a record in the database
func handlePostRequest(w http.ResponseWriter, r *http.Request) (err error) {
	// initialize a byte string of the same length as the input request
	len := r.ContentLength
	body := make([]byte, len)

	// read the input contents into the body variable
	r.Body.Read(body)

	var post Post

	json.Unmarshal(body, &post)

	err = post.create()
	if err != nil {
		fmt.Println("cannot handle POST request and create the record")
		return
	}
	w.WriteHeader(200)
	return
}





// function to handle updates to the posts
func handlePutRequest(w http.ResponseWriter, r *http.Request) (err error){
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		fmt.Println("cannot retrive id of the record that needs to be updated")
		return
	}

	post, err := retrieve(id)

	if err != nil {
		fmt.Println("cannot retrieve the record from database")
		return
	}

	len := r.ContentLength
	body := make([]byte, len)

	r.Body.Read(body)

	json.Unmarshal(body, &post)

	err = post.update()

	if err != nil {
		fmt.Println("cannot update the record in the database")
		return
	}
	w.WriteHeader(200)
	return
}


func handleDeleteRequest(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		fmt.Println("cannot get the data from request. Trying to delete the content")
		return
	}

	post, err := retrieve(id)

	if err != nil{
		fmt.Println("cannot retrieve the data fom database. We are trying to delete the record")
		return
	}

	err = post.delete()
	if err != nil {
		fmt.Println("cannot delete the record from database")
		return
	}

	w.WriteHeader(200)
	return
}
