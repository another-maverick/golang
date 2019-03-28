package main

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

var flagopts struct {
	Name    string `short:"n" long:"name" default:"World!" description:"Who do you want to say Hello to"`
	Spanish bool   `short:"s" long:"spanish" description:"Use Spanish language"`
}

func main() {
	flags.Parse(&flagopts)

	if flagopts.Spanish == true {
		fmt.Printf("Hola  %s \n", flagopts.Name)
	} else {
		fmt.Printf("Hello %s \n", flagopts.Name)
	}
}
