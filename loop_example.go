package main

import "fmt"

func main() {
	for i := 0; i <= 50; i++ {
		fmt.Printf("The outer loop is %d \n", i)
		for j := 0; j <= 30; j++ {

			fmt.Printf("\t \t the inner loop %d \n", j)
			fmt.Println(" just another print stmt", j)
		}
	}
}
