package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "using_cli1"
	app.Usage = "first cut at using the cli right"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Usage: "Who do u want to say Hello",
			Value: "World!!",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		fmt.Printf("Hello %s! \n", name)
		return nil
	}

	app.Run(os.Args)
}
