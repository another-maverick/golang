package main

import "fmt"

func main() {
	x := 5

	fmt.Printf("The type of field is %T, address is %v and value is %v \n", x, &x, x)
}
