package main

import "fmt"

func main() {
	x := 2
	fmt.Println("printing the original value of x in decimal and binary format")
	fmt.Printf("%d \t \t \t %b\n", x, x)

	fmt.Println("Now lets shift the bits")

	y := x << 1

	fmt.Printf("%d \t \t %b\n", y, y)
}
