package main

import (
	"log"
	"testing"
)

var testABData = [][]string{
	{"active", "A"},
	{"inactive", "I"},
	{"banned", "B"},
}

func TestStatusMapper(t *testing.T) {
	for _, d := range testABData {
		s := statusA.Key(cardStatusType).ServiceA(d[0]).ToServiceB()
		log.Printf("a: %s, b: %s", d[0], s)
		if s != d[1] {
			t.Errorf("want %s got %s", d[1], s)
		}
	}
}
