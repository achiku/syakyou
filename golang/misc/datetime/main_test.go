package main

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	y := int64(2016)
	m := int64(8)

	d := time.Date(int(y), time.Month(m), 1, 0, 0, 0, 0, time.Local)
	t.Logf("%s", d)
}
