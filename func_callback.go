package main

import "fmt"

func main() {
	//inp := []int{1, 2, 3, 4, 5}

	g := func(slc []int) int {
		for i := 0; i < len(slc); i++ {

			fmt.Println(i)

		}
		return 6666
	}
	ret := foo(g)
	fmt.Println(ret)
}

func foo(bar func(slc []int) int) int {

	inp := []int{1, 2, 3, 4, 5}
	n := bar(inp)

	return n

}
