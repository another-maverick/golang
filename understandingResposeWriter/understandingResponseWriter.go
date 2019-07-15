package main

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
)

type CustomResponse struct{
	Player string
	Role []string
}


// In this function we demonstrate that responseWriter would set content-type based on the data in response. Test using curl and check response Headers
func normalResponse(w http.ResponseWriter, r *http.Request){
	respStr := `<html>
<head><title> My Custom Response </title> </head>
<body> <h1> Testing the response headers ...... </h1></body>
</html>`
w.Write([]byte(respStr))
}


// sending custom HTTP response code. NOte you cant set headers after setting http response code - precedence wise
func customResponseCode(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(501)
	fmt.Fprintln(w, "You have reached an endpoint that does not exist")

}


// Redirect: set statuscode to 302 and add Location header. Note the precedence of Header set before writeHeader
func testingRedirect(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Location", "http://apple.com")
	w.WriteHeader(302)
}

// JSON response using the CustomResponse struct
func writeJSONResponse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	customResponse := &CustomResponse{
		Player: "Stephen Curry",
		Role: []string{"Guard", "Captain", "Face of the team"},
	}
	json, _ := json2.Marshal(customResponse)
	w.Write(json)
}


func main(){
	server := http.Server {
		Addr: "127.0.0.1:12345",

	}
	http.HandleFunc("/normalresponse", normalResponse)
	http.HandleFunc("/customresponsecode", customResponseCode)
	http.HandleFunc("/testredirect", testingRedirect)
	http.HandleFunc("/jsonresponse", writeJSONResponse)
	server.ListenAndServe()
}