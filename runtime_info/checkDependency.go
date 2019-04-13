package main

import (
	"fmt"
	"os/exec"
)

func main() {
	res := checkDep("ls")
	fmt.Println(res)
}

func checkDep(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		fmt.Println(name + " not present in path")
		return fmt.Errorf("Could not find %s in PATH. Error is %s", name, err)
	}
	return nil
}
