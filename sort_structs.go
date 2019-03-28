package main

import (
	"fmt"
	"sort"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []User

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j, int)     { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {

	u1 := User{
		First:   "Stephen",
		Last:    "Curry",
		Age:     32,
		Sayings: []string{"Best Shooter"},
	}

	u2 := User{
		First:   "Lebron",
		Last:    "James",
		Age:     36,
		Sayings: []string{"Best Player"},
	}

	users := []User{u1, u2}

	sort.Sort(ByAge(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age, u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}

	}

}
