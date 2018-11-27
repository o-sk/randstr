package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"

	"github.com/urfave/cli"
)

var letters = []rune("abcdefghijklmnopgrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func init() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
}

func main() {
	app := cli.NewApp()

	app.Name = "randstr"
	app.Usage = "Return random string"
	app.Version = "0.0.1"

	var length int
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "length, l",
			Usage:       "Length or random string",
			Value:       20,
			Destination: &length,
		},
	}

	app.Action = func(conrext *cli.Context) error {
		string := randstr(length)
		fmt.Println(string)
		return nil
	}

	app.Run(os.Args)
}

func randstr(length int) string {
	runes := make([]rune, length)
	for i := range runes {
		runes[i] = letters[rand.Intn(len(letters))]
	}
	return string(runes)
}
