package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"

	"github.com/atotto/clipboard"
)

type RandStr struct {
	String string
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func init() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
}

func (randstr *RandStr) Generate(length int) {
	runes := make([]rune, length)
	for i := range runes {
		runes[i] = letters[rand.Intn(len(letters))]
	}
	randstr.String = string(runes)
}

func (randstr *RandStr) Output(output bool) {
	if output {
		fmt.Println(randstr.String)
	} else {
		clipboard.WriteAll(randstr.String)
	}
}
