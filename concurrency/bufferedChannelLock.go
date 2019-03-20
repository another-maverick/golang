package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool, 1)

	for i := 0; i <= 10; i++ {
		go worker(i, ch)
	}
	time.Sleep(10 * time.Second)
}

func worker(i int, ch chan bool) {
	fmt.Printf("%d wants to acquire the lock \n", i)
	ch <- true
	fmt.Println("=== Acquired ===  ", i)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%d wants to release the lock \n", i)
	fmt.Println("=== released ===")
	<-ch
}
