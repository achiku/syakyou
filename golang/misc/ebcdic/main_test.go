package main

import (
	"testing"

	"golang.org/x/text/encoding/charmap"
)

func TestHelloWorld(t *testing.T) {
	// EBCDIC decoder
	decoder := charmap.CodePage037.NewDecoder()
	// EBCDIC "LOGN"
	utf8, err := decoder.Bytes([]byte("\xd3\xd6\xc7\xd5"))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", utf8)
}
