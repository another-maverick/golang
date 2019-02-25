package main

import (
	"fmt"
	"sort"
)

func main() {
	sliceint := []int{2, 1, 3, 2, 13, 123, 432, 1231, 55, 66, 77, 44, 33, 22, 11, 12345}
	slicestr := []string{"Hello", "random", "shit", "everywhere", "and", "take", "shit"}

	fmt.Println("unsorted ", sliceint)
	fmt.Println("unsorted str ", slicestr)

	sort.Ints(sliceint)
	sort.Strings(slicestr)

	fmt.Println("sorted --- ", sliceint)
	fmt.Println("sorted str ---", slicestr)

}
