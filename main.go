package main

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"os"

	"github.com/urfave/cli"
)

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
			Usage:       "Length of random string",
			Value:       20,
			Destination: &length,
		},
	}

	app.Action = func(conrext *cli.Context) error {
		var randstr RandStr
		randstr.Generate(length)
		randstr.Output()
		return nil
	}

	app.Run(os.Args)
}
