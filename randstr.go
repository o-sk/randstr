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

var alphabets = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers = []rune("1234567890")
var symbols = []rune("+-*%<=>?@$&_.")

func init() {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
}

func (randstr *RandStr) Generate(length int) {
	randstr.generate(append(append(alphabets, numbers...), symbols...), length)
}

func (randstr *RandStr) GenerateAlphabet(length int) {
	randstr.generate(alphabets, length)
}

func (randstr *RandStr) GenerateNumber(length int) {
	randstr.generate(numbers, length)
}

func (randstr *RandStr) generate(letters []rune, length int) {
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
