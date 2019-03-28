package main

import "fmt"

func main() {
	x := []int{12, 14, 24, 456, 123, 2310}
	x = append(x, 1, 2, 1, 3, 1, 56)
	y := []int{534, 1231, 412, 345}
	x = append(x, y...)
	fmt.Println(x)
}
