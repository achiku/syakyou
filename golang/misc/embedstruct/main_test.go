package main

import "testing"

func TestEmbedAndString(t *testing.T) {
	s := SecuerStruct{
		Name:     "achiku",
		Password: "secret",
	}
	t.Log(s)
}
