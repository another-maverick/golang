package main

import "fmt"

func main() {
	x := "foo bar 1"

	if x == "foo bar" {
		fmt.Printf("Yes, that is correct x is %v \n", x)
	} else {
		fmt.Println("you are not right") // Note that you have to write else in the same line as closing braces
	}

}
