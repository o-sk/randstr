package main

import (
	"os"

	"github.com/o-sk/randstr"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "randstr"
	app.Usage = "Copy random strings to clipboard"
	app.Version = "0.0.1"

	var (
		length int
		output bool
	)

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "length, l",
			Usage:       "Length of random string",
			Value:       20,
			Destination: &length,
		},
		cli.BoolFlag{
			Name:        "o",
			Usage:       "Output stdout",
			Destination: &output,
		},
	}

	app.Action = func(conrext *cli.Context) error {
		var randstr randstr.RandStr
		randstr.Generate(length)
		randstr.Output(output)
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "alphabet",
			Aliases: []string{"a"},
			Usage:   "alphabet only",
			Action: func(context *cli.Context) error {
				var randstr randstr.RandStr
				randstr.GenerateAlphabet(length)
				randstr.Output(output)
				return nil
			},
		},
		{
			Name:    "number",
			Aliases: []string{"n"},
			Usage:   "number only",
			Action: func(context *cli.Context) error {
				var randstr randstr.RandStr
				randstr.GenerateNumber(length)
				randstr.Output(output)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
