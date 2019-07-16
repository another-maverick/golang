package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

 // adding a cookie to the response
func customCookie(w http.ResponseWriter, r *http.Request) {
	cookieData := []byte("Setting data in Cookie!!!!")
	c := http.Cookie{
		Name: "flash",
		Value: base64.URLEncoding.EncodeToString(cookieData),
	}
	http.SetCookie(w, &c)
}


// function to show a temporary message to the user
func showTransientMessage(w http.ResponseWriter, r *http.Request){
	cookieFromRequest, err := r.Cookie("flash")

	// Here we check if the cookie is present in the request. If it is present, set a custom cookie with already expiry time basically, deleting the cookie
	if err != nil{
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "I have set cookie to expire in the past. Hence No cookie found")
		}
	} else {
		tempCookie := http.Cookie{
			Name: "flash",
			MaxAge: -1,
			Expires: time.Unix(1,0),
		}
		http.SetCookie(w, &tempCookie)
		val, _ := base64.URLEncoding.DecodeString(cookieFromRequest.Value)
		fmt.Fprintln(w, string(val))
	}
}


func main() {
	server := http.Server{
		Addr: "127.0.0.1:12345",

	}
	http.HandleFunc("/set_cookie", customCookie)
	http.HandleFunc("/show_and_delete_cookie", showTransientMessage)
	server.ListenAndServe()
}
