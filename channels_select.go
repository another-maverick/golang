package main

import "fmt"

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i <= 100; i++ {
			c <- i
		}
		q <- 10
		close(c)
	}()
	return c
}

func receive(c, q <-chan int) {

	for {
		select {
		case val := <-c:
			fmt.Println(val)
		case <-q:
			return
		}
	}

}
