package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tci"
	app.Usage = "continuous integration in terminal"

	app.Commands = []cli.Command{
		{
			Name:    "deploy",
			Aliases: []string{"d"},
			Usage:   "run deploy script",
			Action:  func(c *cli.Context) error {
				fmt.Println("Deployed!")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}