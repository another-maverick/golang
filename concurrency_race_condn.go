package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementor := 0
	gs := 100

	wg.Add(gs)

	for i := 0; i <= gs; i++ {
		go func() {
			val := incrementor

			runtime.Gosched()
			val++
			incrementor = val
			fmt.Println("from the loop , incrementor = ", incrementor)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("before exiting , incrementor = ", incrementor)
	fmt.Println("exiting")
}
