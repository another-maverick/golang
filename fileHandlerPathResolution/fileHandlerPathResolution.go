package main

import (
	"net/http"
	"path"
	"fmt"
)

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}





func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}

	http.NotFound(res, req)
}

func testFunc(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(res, "Hello, my name is ", name)
}

func main() {
	pr := newPathResolver()

	pr.Add("GET /testPath", testFunc)

	dir := http.Dir("./files")
	handler := http.StripPrefix("/static/", http.FileServer(dir))

	pr.Add("GET /static/*", handler.ServeHTTP)

	http.ListenAndServe(":12345", pr)
}