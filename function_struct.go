package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}

func (p person) speak() {
	fmt.Println("My name is ", p.firstname, p.lastname, " and my age is", p.age)
}

func main() {
	p1 := person{
		firstname: "Stephen",
		lastname:  "Curry",
		age:       31,
	}

	p1.speak()

}
