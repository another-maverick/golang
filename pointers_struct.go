package main

import "fmt"

type person struct {
	firstname string
	level     string
}

func main() {
	p1 := person{
		firstname: "Stephen",
		level:     "good",
	}
	fmt.Println("before changing: ", p1.firstname, " is ", p1.level)

	ChangeLevel(&p1, "Legend")

	fmt.Println("after changing: ", p1.firstname, " is ", p1.level)

}

func ChangeLevel(p *person, newval string) {
	(*p).level = newval
}
