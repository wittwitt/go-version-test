package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "cli",
		Usage:   "cli for",
		Version: "v0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "config.toml",
			},
		},
		Before: func(c *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			subCmd(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
	}
}

func subCmd() *cli.Command {
	return &cli.Command{
		Name: "sub",
		Action: func(ctx *cli.Context) error {
			return nil
		},
	}
}
