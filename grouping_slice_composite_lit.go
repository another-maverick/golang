package main

import "fmt"

func main() {
	x := []int{234, 123, 32432, 21, 31432, 123, 132, 1231, 12412, 122}

	for i, v := range x {
		fmt.Println("Hello")
		fmt.Printf("value:%v, Type:%T, index:%d \n", v, v, i)
	}

}
