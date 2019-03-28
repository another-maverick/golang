package main

import "fmt"

func main() {
	fmt.Println(true && true)
	fmt.Println(false && false)
	fmt.Println(false && true)
	fmt.Println(true && false)
	fmt.Println(!true)
	fmt.Println(true || true)
	fmt.Println(false || false)
	fmt.Println(false || true)
	fmt.Println(true || false)
}
