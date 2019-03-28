package main

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config.Get("mypath"))
	fmt.Println(config.Get("enabled"))
}
