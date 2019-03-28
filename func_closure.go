package main

import "fmt"

func main() {
	res := foo()

	res1 := res()
	fmt.Println(res1)
}

func foo() func() int {

	x := 0
	return func() int {
		x++
		return x
	}
}
