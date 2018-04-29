package main

import (
	"encoding/json"
	"testing"
)

func TestEncode(t *testing.T) {
	h := hoge{
		Val: "test",
		Num: 10,
	}
	b, err := json.Marshal(h)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", b)
}
