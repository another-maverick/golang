package main

import "fmt"

func main() {
	fmt.Println("Hello, all we are learning Go Lang")
	foo()
	fmt.Println("doing something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}

func foo() {
	fmt.Println("I am in foo")
}

//control flow
// sequence
// loop
// conditional
