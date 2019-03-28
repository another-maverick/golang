package main

import (
	"fmt"

	gcfg "gopkg.in/gcfg.v1"
)

func main() {

	config := struct {
		Myarea struct {
			Enabled bool
			Mypath  string
		}
	}{}
	err := gcfg.ReadFileInto(&config, "config.ini")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Myarea.Enabled)
	fmt.Println(config.Myarea.Mypath)
}
