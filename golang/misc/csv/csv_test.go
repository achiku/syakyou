package main

import (
	"encoding/csv"
	"os"
	"strings"
	"testing"
)

var msg = "123|abcd|dddd|08/19||"

func TestTempFile(t *testing.T) {
	f, err := os.Open("test.csv")
	if err != nil {
		t.Fatal(err)
	}
	r := csv.NewReader(f)
	r.Comma = '|'
	recs, err := r.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	for _, r := range recs {
		t.Logf("123=%s", r[0])
		t.Logf("abcd=%s", r[1])
		t.Logf("dddd=%s", r[2])
		t.Logf("08/19=%s", r[3])
		t.Logf("=%s", r[4])
		t.Logf("=%s", r[5])
	}
}

func TestReadXSeparatedFile(t *testing.T) {
	r := csv.NewReader(strings.NewReader(msg))
	r.Comma = '|'
	recs, err := r.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	for _, r := range recs {
		t.Logf("123=%s", r[0])
		t.Logf("abcd=%s", r[1])
		t.Logf("dddd=%s", r[2])
		t.Logf("08/19=%s", r[3])
		t.Logf("=%s", r[4])
		t.Logf("=%s", r[5])
	}
}
