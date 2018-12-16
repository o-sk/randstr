package main

import (
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rand.Seed(123456)
	os.Exit(m.Run())
}
