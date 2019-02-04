package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

type details struct {
	team  string
	years int
}

type profile struct {
	player      person
	information details
}

func main() {
	player1 := profile{
		player: person{
			firstname: "Stephen",
			lastname:  "Curry",
		},
		information: details{
			team:  "warriors",
			years: 7,
		},
	}
	fmt.Println(player1)
	fmt.Println(player1.player)
	fmt.Println(player1.player.firstname)
	fmt.Println(player1.information.team)
}
