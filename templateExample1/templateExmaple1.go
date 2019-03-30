package main

import (
	"html/template"
	"net/http"
)

type TemplateTest struct {
	Title string
	Content string
}

func myPage(res http.ResponseWriter, req *http.Request) {
	tTest := &TemplateTest {
		Title: "Champions",
		Content: "Waaaaaaaaaarriors!!!!",
	}

	t := template.Must(template.ParseFiles("test.html"))
	t.Execute(res, tTest)


}

func main() {
	http.HandleFunc("/hello", myPage)
	http.ListenAndServe(":12345", nil)
}