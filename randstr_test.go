package main

import "testing"

func TestGeneraate(t *testing.T) {
	var randstr RandStr
	randstr.Generate(20)

	expect := "6SgUn6wO2Df2LXVmzZOs"
	if randstr.String != expect {
		t.Errorf("Expect randstr.String is %s, but %s", expect, randstr.String)
	}
}
