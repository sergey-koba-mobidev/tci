package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/sergey-koba-mobidev/tci/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "tci"
	app.Usage = "continuous integration in terminal"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Serhii Koba",
			Email: "s.koba@mobidev.biz",
		},
	}

	var filename string

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "file, f",
			Value: "tci.yml",
			Usage: "yml file with deploy steps",
			Destination: &filename,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "deploy",
			Aliases: []string{"d"},
			Usage:   "run deploy script",
			Action:  func(c *cli.Context) error {
				err := commands.Deploy(filename)
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}