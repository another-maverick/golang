package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {

	u1 := User{
		First:   "Stephen",
		Last:    "Curry",
		Age:     32,
		Sayings: []string{"Best shooter", "champion"},
	}
	u2 := User{
		First:   "Lebron",
		Last:    "James",
		Age:     35,
		Sayings: []string{"Best player", "champion"},
	}

	users := []User{u1, u2}
	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}

}
