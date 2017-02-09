package main

import "testing"

func TestGenerateCode(t *testing.T) {
	code := generateCode()
	t.Logf("%s", code)
}
