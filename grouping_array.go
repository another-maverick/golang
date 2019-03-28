package main

import "fmt"

func main() {
	x := [5]int{1, 3, 45, 56, 13}

	for i, v := range x {
		fmt.Printf("the current element value is %v and type is %T and index is %d\n", v, v, i)
	}
}
