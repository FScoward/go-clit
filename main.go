package main

import (
	"os"
	"github.com/urfave/cli"
	"fmt"
)

func main() {
	app := cli.NewApp()

	app.Name = "go clit"
	app.Usage = "Task Manage"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
			{
				Name: "new",
				Usage: "create",
				Aliases: []string{"n"},
				Subcommands: []cli.Command{
					{
						Name:"project",
						Usage: "project",
						Action: func(c *cli.Context) error {
							fmt.Println(c.Args().First())
							return nil
						},
					},
					{
						Name:"task",
						Usage: "task",
						Action: func(c *cli.Context) error {
							if c.NArg() != 1 {
							}

							fmt.Println(c.Args().First())
							return nil
						},
					},
				},
			},
		}

	app.Run(os.Args)
}
