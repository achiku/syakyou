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
