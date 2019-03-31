package main

import "net/http"


func myFunc(res http.ResponseWriter, req *http.Request){

	http.ServeFile(res, req, "./files/readme.txt")

}


func main() {
	dir := http.Dir("./files/")

	handler := http.StripPrefix("/static/", http.FileServer(dir))
	http.Handle("/static/", handler)

	http.HandleFunc("/", myFunc)
	http.ListenAndServe(":12345", nil)
}