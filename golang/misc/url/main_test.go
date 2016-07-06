package main

import (
	"net/url"
	"testing"
)

func TestURLQuery(t *testing.T) {
	data := []string{
		"key1=value1&key2=value2",
		"value1&key2=value2",
	}
	for _, d := range data {
		s, err := url.ParseQuery(d)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(s)
	}
}
