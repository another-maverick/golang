package main

import "fmt"

func main() {
	m := map[string]int{
		"James":     32,
		"other guy": 45,
	} // key is a string, value is an int
	fmt.Println(m)

	v, ok := m["random"]
	fmt.Println(v)
	fmt.Println(ok)
	// get value and ok(true or false) and then in the same line check if ok is true, then execute the loop

	if v, ok := m["James"]; ok {
		fmt.Println("I am  present", v)
	}

	m["new person"] = 23

	for k, v := range m {
		fmt.Println(k, "------", v)
	}

}
