package main

import "fmt"

var x = 1235

var y = "hello there"

var z = 123.456

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
