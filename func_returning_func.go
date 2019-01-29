package main

import "fmt"

func main() {

	p1 := foo()
	fmt.Println(p1)

	p2 := bar()
	i := p2()
	fmt.Println(i)
	fmt.Printf("%T\n", p2)

}

func foo() string {
	s := "Hello World"
	return s
}

func bar() func() int {
	return func() int {
		return 1234
	}
}
