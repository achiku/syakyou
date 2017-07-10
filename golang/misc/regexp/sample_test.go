package main

import "testing"

func TestRegexp(t *testing.T) {
	cases := []string{
		"aaaaa_b bb:",
	}

	for _, c := range cases {
		s := r.ReplaceAllString(c, "$1")
		t.Logf("str->%s", s)
	}
}
