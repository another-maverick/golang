package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "Stephen",
		Last:  "Curry",
		Age:   32,
	}
	json_bs, err := json.Marshal(p1)

	if err != nil {
		log.Fatalln("error is -- ", err)
	}

	fmt.Println("output is ---- ", json_bs)
	fmt.Println("output in string format is ---", string(json_bs))

}
