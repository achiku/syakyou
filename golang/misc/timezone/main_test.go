package main

import (
	"testing"
	"time"
)

func TestLocalTimezone(t *testing.T) {
	tm := time.Now()
	tmUTC := tm.UTC()
	t.Logf("%s", tm)
	t.Logf("%s", tmUTC)
}

func TestHandlerUTCDateCol(t *testing.T) {
	JST := time.FixedZone("JST", 9*60*60)
	tm := time.Date(2017, 8, 18, 0, 0, 0, 0, time.UTC)
	tmJST := tm.In(JST)
	t.Logf("%s", tm)
	t.Logf("%s", tmJST)
	t.Logf("%s", tm.Sub(tmJST))
}
