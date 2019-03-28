package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("switch is empty or true, so this should print")
	case false:
		fmt.Println("switch is false")
	}

}
