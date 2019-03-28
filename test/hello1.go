package main

import "fmt"

func main() {
	name := foo()
	fmt.Println("Hello ", name)
}

func foo() string {
	return "World!!"
}
