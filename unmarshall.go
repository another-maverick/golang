package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {

	s := `[{"First":"Stephen", "Last":"Curry", "Age":32, "Sayings":["Best shooter"]}, {"First":"Lebron","last":"james", "Age":35, "Sayings":["Best player"]}]`

	fmt.Println(s)

	var people []person

	err := json.Unmarshal([]byte(s), &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, plyr := range people {
		fmt.Println("player #", i)
		fmt.Println("\t", plyr.First, plyr.Last, plyr.Age)
		for _, v := range plyr.Sayings {
			fmt.Println("\t\t", v)
		}
	}

}
