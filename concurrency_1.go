package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begenning CPU's", runtime.NumCPU())
	fmt.Println("begenning routines", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("hello from first routine")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from second routine")
		wg.Done()
	}()

	fmt.Println("mid CPU's", runtime.NumCPU())
	fmt.Println("mid routines", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("before exiting")
	fmt.Println("ending CPU's", runtime.NumCPU())
	fmt.Println("ending routines", runtime.NumGoroutine())

}
