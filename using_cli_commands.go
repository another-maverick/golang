package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {

	app := cli.NewApp()
	app.Name = "using_cli_commands"
	app.Usage = "Using global flags commands and sub commands"

	app.Commands = []cli.Command{
		{
			Name:      "up",
			Usage:     "Count UP",
			ShortName: "u",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "stop, s",
					Usage: "Value that you can count upto",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("stop")
				if start <= 0 {
					fmt.Println("You have reached zero")
				}
				for i := 1; i <= start; i++ {
					fmt.Println(i)
				}
				return nil
			},
		},
		{
			Name:      "down",
			Usage:     "Count Down",
			ShortName: "d",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start, s",
					Usage: "Value that can count down from",
					Value: 20,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("start")
				if start < 0 {
					fmt.Println("you cant start from negative")
				}
				for i := start; i >= 0; i-- {
					fmt.Println(i)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
