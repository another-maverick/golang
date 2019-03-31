package main

import ("fmt"
	myfs "github.com/Masterminds/go-fileserver"
	"net/http"
)

func main() {
	myfs.NotFoundHandler = func(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(res, "This is my custom 404 page")
	}
	dir := http.Dir("./somerandomdir")
	http.ListenAndServe(":12345", myfs.FileServer(dir))
}

	
