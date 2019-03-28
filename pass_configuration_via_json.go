package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, _ := os.Open("conf_file.json")
	defer file.Close()

	mydecoder := json.NewDecoder(file)
	conf := configuration{}
	err := mydecoder.Decode(&conf)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("some data from the decoded json --- ", conf.Path)

}
