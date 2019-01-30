package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x, before calling the foo function")
	fmt.Println(x)
	foo(x)
	fmt.Println("x, after calling the foo function")
	fmt.Println(x)
	bar(&x)
	fmt.Println("x, after calling the bar function")
	fmt.Println(x)
}

func foo(y int) {
	fmt.Println(y)
	y = 45
	fmt.Println(y)
}

func bar(z *int) {
	fmt.Println(z)
	*z = 456
	fmt.Println(z)
}
