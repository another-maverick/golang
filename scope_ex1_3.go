package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%T\t%T\t%T\n", x, y, z)
	s1 := fmt.Sprintf("%v\t%v\t%v\n", x, y, z)
	fmt.Println(s)
	fmt.Println(s1)
}
