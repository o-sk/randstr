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

func TestGeneraateAlphabet(t *testing.T) {
	var randstr RandStr
	randstr.GenerateAlphabet(20)

	expect := "KXbRMJaEefpoOQDKRlbO"
	if randstr.String != expect {
		t.Errorf("Expect randstr.String is %s, but %s", expect, randstr.String)
	}
}

func TestGeneraateNumber(t *testing.T) {
	var randstr RandStr
	randstr.GenerateNumber(8)

	expect := "23843497"
	if randstr.String != expect {
		t.Errorf("Expect randstr.String is %s, but %s", expect, randstr.String)
	}
}
