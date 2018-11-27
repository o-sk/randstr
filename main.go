package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/urfave/cli"
)

var letters = []rune("abcdefghigeklmnopgrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

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
	runes := make([]rune, 8)
	for i := range runes {
		runes[i] = letters[rand.Intn(len(letters))]
	}
	return string(runes)
}
