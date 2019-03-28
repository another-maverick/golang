package main

import (
	"encoding/json"
	"fmt"
)

type people []struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	mystr := `[{"First":"Rakesh", "Last":"Vad", "Age": 32}, {"First":"Virat", "Last":"kohli", "Age": 31}]`
	mybytestr := []byte(mystr)
	fmt.Printf("%T\n", mystr)
	fmt.Printf("%T\n", mybytestr)

	var mypeople people

	err := json.Unmarshal(mybytestr, &mypeople)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("printing people data", mypeople)
}
