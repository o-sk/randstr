package main

import "testing"

func TestGenerate(t *testing.T) {
	var randstr RandStr
	randstr.Generate(20)

	expect := "I>7ZIMCm@Fn2_gbk_?gt"
	if randstr.String != expect {
		t.Errorf("Expect randstr.String is %s, but %s", expect, randstr.String)
	}
}

func TestGenerateAlphabet(t *testing.T) {
	var randstr RandStr
	randstr.GenerateAlphabet(20)

	expect := "KXbRMJaEefpoOQDKRlbO"
	if randstr.String != expect {
		t.Errorf("Expect randstr.String is %s, but %s", expect, randstr.String)
	}
}

func TestGenerateNumber(t *testing.T) {
	var randstr RandStr
	randstr.GenerateNumber(8)

	expect := "23843497"
	if randstr.String != expect {
		t.Errorf("Expect randstr.String is %s, but %s", expect, randstr.String)
	}
}
