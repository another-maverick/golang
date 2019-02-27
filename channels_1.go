package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		c <- 123
	}()

	fmt.Println(<-c)
}
