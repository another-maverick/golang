package main

import "fmt"

func main() {
	x := []int{23, 33, 56, 76, 190}
	for i, v := range x {
		fmt.Println(i, "=====", v)
	}
}
