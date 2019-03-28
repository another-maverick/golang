package main

import "fmt"

func main() {
	inp := []int{1, 2, 34, 5, 6, 7, 8657, 234, 112, 33}
	res := foo(inp...)
	fmt.Println("result obtained is ", res)
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v

	}
	return total
}
