package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 123, 4, 3243, 32432}

	fmt.Println("original", x)

	fmt.Println("now slicing")

	x = append(x[:2], x[5:]...)
	fmt.Println("afterwards", x)
}
