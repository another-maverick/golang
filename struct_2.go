package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	sports    []string
}

func main() {

	person1 := person{
		firstname: "james",
		lastname:  "bond",
		sports:    []string{"football", "volleyball"},
	}

	person2 := person{
		firstname: "lebron",
		lastname:  "james",
		sports:    []string{"basketball", "soccer"},
	}
	fmt.Println(person1)
	fmt.Println(person2)

	for i, v := range person1.sports {
		fmt.Println(i, v)
	}

	mymap := map[string]person{person1.firstname: person1, person2.firstname: person2}

	for k, v := range mymap {
		fmt.Println(k)
		fmt.Println(v)
	}
}
