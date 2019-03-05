package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World!", "A random name to say hello")
var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish Language")
	flag.BoolVar(&spanish, "s", false, "Use spanish language")
}
func main() {
	flag.Parse()

	flag.VisitAll(func(flag *flag.Flag) {
		format := "\t %s - %s (Default: '%s')\n"
		fmt.Printf(format, flag.Name, flag.Usage, flag.Value)
	})

	if spanish == true {
		fmt.Printf("Hola %s\n", *name)
	} else {
		fmt.Printf("Hello %s\n", *name)
	}
}
