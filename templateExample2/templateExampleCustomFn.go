package main

import (
	"html/template"
	"net/http"
	"time"
)

var temp1 = ` <!DOCTYPE HTML>
<html>
<head>
<meta charset="utf-8">
<title> 
Template with Custom Function
</title>
</head>
<body>
<p>
{{ .Date |dateFormat "Mar 29, 2019"}}
</p>
</body>
</html>`


var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

func dateFormat(custFormat string, f1 time.Time) string {
return f1.Format(custFormat)

}

func tempTestWithFn(res http.ResponseWriter, req *http.Request) {
	t := template.New("test")
	t.Funcs(funcMap)
	t.Parse(temp1)

	dataInTemplate := struct{ Date time.Time} {
		Date: time.Now(),
	}
	t.Execute(res, dataInTemplate)
}

func main() {
	http.HandleFunc("/customfunction", tempTestWithFn)
	http.ListenAndServe(":12345", nil)
}

