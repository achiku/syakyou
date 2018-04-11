package main

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHexAndString(t *testing.T) {
	str := "1"
	x := make([]byte, hex.EncodedLen(len(str)))
	hex.Decode(x, []byte(str))
	t.Log(x)
	t.Logf("%x", x)
}

func TestHoge(t *testing.T) {
	s := "9161"
	decoded, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v", decoded)
	fmt.Printf("%x", decoded)
}
