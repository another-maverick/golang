package main

import (
	"fmt"
	"time"
)

//var wg sync.WaitGroup

func foo(ch chan int) {
	num := 0
	for num >= 0 {
		//defer wg.Done()

		num = <-ch
		fmt.Println("  current number from channel --- ", num)
	}

}

func main() {

	a := []int{1, 2, 24, 5, 12, 1, 313, 12, 43, 21, 44, 67, 452, 113, 45, 01, 33, -11}
	ch := make(chan int)
	//wg.Add(1)
	go foo(ch)

	for _, v := range a {
		ch <- v
	}
	time.Sleep(time.Millisecond * 1000)

	fmt.Println("end of main")

}
