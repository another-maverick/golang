package main

import "fmt"

var x int

// define your own type. here hotdog is of type int
type hotdog int

var y hotdog

func main() {
	x = 123
	y = 3456
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println("now printing Y and type of Y")
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
