package main

import (
	"strconv"
	"strings"
	"testing"
)

type testData struct {
	In  string
	Out string
}

func TestGetLastFour(t *testing.T) {
	data := []testData{
		{"hello, world!", "rld!"},
		{"1234567890", "7890"},
		{"098", "098"},
		{"", ""},
	}

	for _, d := range data {
		s := getLastFour(d.In)
		if s != d.Out {
			t.Errorf("want %s got %s", d.Out, s)
		}
	}
}

func TestGetLastFour2(t *testing.T) {
	data := []testData{
		{"hello, world!", "rld!"},
		{"1234567890", "7890"},
		{"098", "s is less than four chars"},
		{"", "s is less than four chars"},
	}

	for _, d := range data {
		s, err := getLastFour2(d.In)
		if err != nil {
			if d.Out != err.Error() {
				t.Error(err)
			}
		} else {
			if s != d.Out {
				t.Errorf("want %s got %s", d.Out, s)
			}
		}
	}
}

func TestTremSpace(t *testing.T) {
	data := map[string]string{
		"test   ":     "test",
		"   test    ": "test",
		"   test":     "test",
		"  　　test":    "test",
		"  --test":    "--test",
	}

	for k, v := range data {
		s := strings.TrimSpace(k)
		if s != v {
			t.Errorf("want %s got %s", v, s)
		}
	}
}

func TestParseFloat(t *testing.T) {
	n := "1000.01111234445"
	f, err := strconv.ParseFloat(n, 64)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%f", f)
}
