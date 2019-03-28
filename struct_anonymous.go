package main

import "fmt"

func main() {
	anon_s := struct {
		firstname  string
		lastname   string
		mvp_awards []int
	}{ // Note that you need to get both the braces at the same line
		firstname:  "Stephen",
		lastname:   "Curry",
		mvp_awards: []int{2015, 2016},
	}
	fmt.Println(anon_s)
}
