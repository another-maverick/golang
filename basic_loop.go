package main

import "fmt"

func main() {
	birth_day := 1970
	for birth_day <= 2019 {
		fmt.Println(birth_day)
		birth_day += 1
	}
}
