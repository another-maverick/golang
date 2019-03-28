package main

import "fmt"

func main() {
	mysport := "football"
	switch mysport {
	case "cricket":
		fmt.Printf("from cricket: my favoutite sport is %v \n", mysport)
	case "football":
		fmt.Printf("from football: my favoutite sport is %v \n", mysport)
	}
}
