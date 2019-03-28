package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	first string
	age   int
}

func main() {

	u1 := user{
		first: "Stephen",
		age:   32,
	}

	u2 := user{
		first: "Lebron",
		age:   35,
	}

	users := []user{u1, u2}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}
