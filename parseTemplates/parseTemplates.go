package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func templateParse(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("testTemplate.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:12345",
	}
	http.HandleFunc("/templateParse", templateParse )
	server.ListenAndServe()
}


