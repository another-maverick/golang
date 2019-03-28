package main

import "fmt"

func main() {
	x := []string{"james", "bond", "chocloate", "martini"}
	y := []string{"miss", "moneypenny", "strawberry", "hazelnut"}

	mdx := [][]string{x, y}
	fmt.Println(mdx)
}
