package main

import "fmt"

func main() {

	mymap := map[string][]string{"key1": []string{"Rakesh", "Vadlakunta"}, "key2": []string{"james", "bond"}}

	fmt.Println(mymap)

	for k, v := range mymap {
		fmt.Println("current key is", k)
		for k1, v1 := range v {
			fmt.Printf("values for %v and index:%d is  %v\n", k, k1, v1)
		}
	}
}
