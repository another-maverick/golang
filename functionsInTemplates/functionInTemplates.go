package main

import (
	"html/template"
	"net/http"
	"time"
)
//Jan 2 15:04:05 2006 MST
//1   2  3  4  5    6  -7
func funcInTemplate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func parseTemplate(w http.ResponseWriter, r *http.Request){
	funcMap := template.FuncMap{"cleanDate": funcInTemplate }
	t := template.New("testTemplate.html").Funcs(funcMap)
	t, _ = t.ParseFiles("testTemplate.html")
	t.Execute(w, time.Now())

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:12345",
	}
	http.HandleFunc("/parsetemplate", parseTemplate)
	server.ListenAndServe()
}
