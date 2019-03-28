package main

import "fmt"

func main() {

	ret1 := foo()
	ret2_1, ret2_2 := bar()

	fmt.Println("return of func foo is ", ret1)
	fmt.Println("return of func bar is ", ret2_1, ret2_2)
}

func foo() int {
	return 123
}

func bar() (int, string) {
	return 1986, "when I happened"
}
