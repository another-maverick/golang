package main

import (
	"fmt"
	"time"
)

func main() {

	msg := make(chan string)
	done := make(chan bool)
	until := time.After(30 * time.Second)

	go send(msg, done)

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true
			time.Sleep(1 * time.Second)
			return
		}
	}
}

func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			close(ch)
			return
		default:
			ch <- "hellooooo"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
