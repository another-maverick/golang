package main

import(
	"bytes"
	"html/template"
	"net/http"
)

var t *template.Template
var tObj template.HTML

func init(){
	t = template.Must(template.ParseFiles("index.html", "quote.html"))
}

type mainData struct {
	Title string
	Content template.HTML
}

type subData struct {
	Quote string
	Person string
}

func main() {

	sd := &subData {
		Quote: "This is being populated in the html object",
		Person: "Steph Curry",
	}

	var buf bytes.Buffer
	t.ExecuteTemplate(&buf,"quote.html", sd)
	tObj = template.HTML(buf.String())
	http.HandleFunc("/templatehtml", templateHandler)
	http.ListenAndServe(":12345", nil)
}


func templateHandler(res http.ResponseWriter, req *http.Request){
	md := &mainData{
		Title: "Title of template with HTML object",
		Content: tObj,
	}
	t.ExecuteTemplate(res, "index.html", md)
}