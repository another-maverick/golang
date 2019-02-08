package main

import "fmt"

func main() {
	x := 123
	fmt.Println("printing x which occurs before defer function ", x)

	defer defer_func()

	y := "after defer"

	fmt.Println("called after defer ---   but still get printed before defer", y)
}

func defer_func() {
	defer func() {
		fmt.Println("----- deferred from the deferred function ----- and i am anonymous ----")
	}()
	fmt.Println("i am being printed from the defer function")
}
