package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func HasTimedOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}

	}
	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}

func main() {
	file, err := os.Create("file.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	location := "https://example.com/file.zip"
	err = download(location, file, 20)
	if err != nil {
		fmt.Println(err)
		return
	}
	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("I have downloaded %v bytes", fi.Size())

}

func download(location string, file *os.File, retries int64) error {
	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		fmt.Println(err)
	}

	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	current_size := fi.Size()
	if current_size > 0 {
		start_index := strconv.FormatInt(current_size, 10)
		req.Header.Set("Range", "bytes="+start_index+"-")
	}
	cc := &http.Client{Timeout: 5 * time.Minute}
	res, err := cc.Do(req)
	if err != nil && HasTimedOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("unsuccesfull HTTP request: %s", res.Status)
	}

	if res.Header.Get("Accept-Ranges") != "bytes" {
		retries = 0
	}
	_, err = io.Copy(file, res.Body)
	if err != nil && HasTimedOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}
	return nil

}
