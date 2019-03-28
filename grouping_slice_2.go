package main

import "fmt"

func main() {
	x := []int{23, 13, 14, 15, 15, 141, 123, 1234, 232, 454, 21343}

	fmt.Println(x[:])
	fmt.Println(x[:5])
	fmt.Println(x[3:])
	fmt.Println(x[1:5])
	//fmt.Println(x[:-1])
	fmt.Println(x[:len(x)])

}
