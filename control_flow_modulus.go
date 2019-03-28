package main

import "fmt"

func main() {
	//limit := 100

	for i := 0; i <= 100; i++ {
		fmt.Printf("when %v is divided by 4, the reminder is %v \n", i, i%4)
		if i%4 == 0 {
			fmt.Println(i)
		}
	}
}
