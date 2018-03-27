package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"

	"github.com/micahlee/big-earl-cli/command/grow"
	"github.com/micahlee/big-earl-cli/command/preview"
	"github.com/micahlee/big-earl-cli/command/shrink"
)

func main() {
	app := cli.NewApp()

	app.Name = "Big Earl's CLI"
	app.Usage = "interact with Big Earl's Lean URLs API"

	// Global flags for Big Earl
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "big_earl_api",
			Usage:  "api server to use",
			EnvVar: "BIG_EARL_API",
			Value:  "https://big-earl.herokuapp.com/",
		},
	}

	app.Commands = []cli.Command{
		grow.Command,
		preview.Command,
		shrink.Command,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
