package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 123, 213, 412, 43}

	fmt.Println(x)
	x = append(x, 2132)
	fmt.Printf("new x: %v \n", x)
	y := []int{876, 76, 87, 43, 3123}
	x = append(x, y...) // y... meaning we are unfurling.
	fmt.Printf("after extending x:%v \n", x)
}
