package main

import "testing"

func TestFind(t *testing.T) {
	cases := []struct {
		number   string
		expected string
	}{
		{number: "4355667891234567", expected: "visa"},
		{number: "5355667891234567", expected: "master"},
		{number: "5355667891234567", expected: "master"},
		{number: "1355667891234567", expected: "unknown"},
		{number: "3455667891234567", expected: "amex"},
		{number: "3755667891234567", expected: "amex"},
	}

	for _, d := range cases {
		b := find(d.number)
		if b.Type != d.expected {
			t.Errorf("want %s got %s", d.expected, b.Type)
		}
	}
}
