package main

import "fmt"

func main() {
	c := make(chan int)

	for outer := 0; outer <= 20; outer++ {
		go func() {
			for i := 0; i <= 30; i++ {
				c <- i
			}

		}()
	}

	for j := 0; j <= 650; j++ {
		fmt.Println(j, <-c)
	}
}
