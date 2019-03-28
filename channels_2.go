package main

import "fmt"

func main() {
	cs := make(chan int)

	go func() {
		cs <- 1234
	}()
	fmt.Println(<-cs)

}
