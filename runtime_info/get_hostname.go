package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name, err := os.Hostname()

	if err != nil {
		fmt.Println("cannot retrieve hostname")
		return
	}
	addr, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("cannot retrieve IP addresses")
		return
	}

	for _, a := range addr {
		fmt.Println(a)
	}

}
