package main

import (
	"fmt"
	"math/rand"
)

type RandStr struct {
	String string
}

var letters = []rune("abcdefghijklmnopgrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func (randstr *RandStr) Generate(length int) {
	runes := make([]rune, length)
	for i := range runes {
		runes[i] = letters[rand.Intn(len(letters))]
	}
	randstr.String = string(runes)
}

func (randstr *RandStr) Output() {
	fmt.Println(randstr.String)
}
