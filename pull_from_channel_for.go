package main

import "fmt"

func main() {
	c := gen()
	receive(c)

	fmt.Println("before exiting")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for val := range c {
		fmt.Println(val)
	}
}
