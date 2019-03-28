package main

import "fmt"

func main() {
	x1 := []string{"Hello", "this", "is", "the", "first", "string"}
	x2 := []string{"hello", "this", "is", "is", "the", "second", "string"}

	x := [][]string{x1, x2}
	fmt.Println(x)

	for i, v := range x {
		fmt.Printf("slice Number %d in 2-D array \n", i)
		for _, v1 := range v {
			fmt.Println(v1)
		}
	}
}
