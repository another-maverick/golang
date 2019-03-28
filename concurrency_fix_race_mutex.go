package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementor := 0
	gs := 100

	wg.Add(gs)

	for i := 0; i <= gs; i++ {
		go func() {
			var m sync.Mutex
			m.Lock()
			val := incrementor
			val++
			incrementor = val
			fmt.Println("from the loop , incrementor = ", incrementor)
			wg.Done()
			m.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("before exiting , incrementor = ", incrementor)
	fmt.Println("exiting")
}
