package main

import "fmt"

type CustomErr struct {
	Info string
}

func (ce CustomErr) Error() string {

	return fmt.Sprintf("Here is the error -- %v", ce.Info)
}

func main() {

	c1 := CustomErr{
		Info: "I am the custom Error",
	}
	foo(c1)

}

func foo(e error) {

	fmt.Println("foo ran - %v", e)
	//assertion
	fmt.Println("foo ran and here i am asserting info from e - ", e.(CustomErr).Info)
}
