package main


import "net/http"


func fileFunc(res http.ResponseWriter, req *http.Request){

	http.ServeFile(res, req, "./files/readme.txt")
}

func main() {

	http.HandleFunc("/", fileFunc)
	http.ListenAndServe(":12345", nil)
}