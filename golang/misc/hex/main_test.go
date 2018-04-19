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
	fmt.Printf("%#v\n", decoded)
	fmt.Printf("%x\n", decoded)
}

func TestAppend(t *testing.T) {
	s := "122233"
	decoded, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	a := []byte{0x91, 0x61}
	a = append(a, decoded...)
	fmt.Printf("%#v\n", a)
}
