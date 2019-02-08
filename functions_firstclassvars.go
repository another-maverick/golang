package main

import "fmt"

func main() {
	f := func() {
		for i := 0; i <= 100; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Println("done")
}
