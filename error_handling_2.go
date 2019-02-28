package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		First: "LeBron",
		Last:  "James",
		Age:   32,
	}
	json_bs, err := toJSON(p1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(json_bs)
	fmt.Println(string(json_bs))

}

func toJSON(a interface{}) ([]byte, error) {

	bs, err := json.Marshal(a)

	if err != nil {
		return []byte{}, fmt.Errorf("there was an error while performing toJSON. Error is - %v", err)
	}
	return bs, nil
}
