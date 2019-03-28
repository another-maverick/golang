package main

import "fmt"

const (
	//kb := 1024
	//mb := kb * 1024
	//gb := mb * 1024

	_  = iota
	kb = 1 << (iota * 10) // here iota is 1
	mb = 1 << (iota * 10) // here iota is 2 . so this equals shifting the bit by 20
	gb = 1 << (iota * 10) // here iota = 3
)

func main() {
	fmt.Printf("%d \t\t %b \n", kb, kb)
	fmt.Printf("%d \t\t %b \n", mb, mb)
	fmt.Printf("%d \t\t %b \n", gb, gb)
}
