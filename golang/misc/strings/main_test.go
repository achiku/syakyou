package main

import "testing"

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
