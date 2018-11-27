package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "randstr"
	app.Usage = "Return random string"
	app.Version = "0.0.1"

	app.Action = func(conrext *cli.Context) error {
		string := randstr()
		fmt.Println(string)
		return nil
	}

	app.Run(os.Args)
}

func randstr() string {
	return "abcdefghi"
}
