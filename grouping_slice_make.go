package main

import "fmt"

func main() {
	x := make([]int, 4, 5)

	x = []int{1, 2, 23, 4}

	for i := 0; i < cap(x); i++ {
		fmt.Println(x[i])
	}

}
