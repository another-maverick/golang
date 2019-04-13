package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go startCollectingStats()

	i := 0

	for i < 40 {
		go func() {
			time.Sleep(20 * time.Second)
		}()
		i += 1
		time.Sleep(1 * time.Second)
	}
}

func startCollectingStats() {
	fmt.Println("number of CPU's ", runtime.NumCPU())

	m := &runtime.MemStats{}
	for {
		r := runtime.NumGoroutine()
		fmt.Println("number of go routines now is ", r)

		runtime.ReadMemStats(m)
		fmt.Println("currently allocated memory is ", m.Alloc)

		time.Sleep(10 * time.Second)
	}
}
