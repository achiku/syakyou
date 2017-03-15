package main

import "testing"

func TestGetSortedStringKeys(t *testing.T) {
	data := map[string]string{
		"apple":  "1000",
		"orange": "100",
		"rice":   "1000",
	}

	keys := getSortedStringKeys(data)
	t.Logf("%s", keys)
}
