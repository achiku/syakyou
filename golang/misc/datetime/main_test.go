package main

import (
	"math"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	y := int64(2016)
	m := int64(8)

	d := time.Date(int(y), time.Month(m), 1, 0, 0, 0, 0, time.Local)
	t.Logf("%s", d)
}

func TestInterval(t *testing.T) {
	// start
	startDate := time.Date(2016, 9, 16, 0, 0, 0, 0, time.Local)
	// 120 after start
	targetDate1 := time.Date(2017, 1, 14, 0, 0, 0, 0, time.Local)
	// 150 after start
	targetDate2 := time.Date(2017, 2, 13, 0, 0, 0, 0, time.Local)
	// 180 after start
	targetDate3 := time.Date(2017, 3, 15, 0, 0, 0, 0, time.Local)

	// other dates
	anotherDate1 := time.Date(2017, 1, 15, 0, 0, 0, 0, time.Local)
	anotherDate2 := time.Date(2017, 1, 22, 0, 0, 0, 0, time.Local)
	anotherDate3 := time.Date(2017, 3, 22, 0, 0, 0, 0, time.Local)

	duration := float64(2880)
	thirtyDaysCycle := float64(720)
	// 24 * 120 = 2880
	// 24 * 30 = 720

	diffHr1 := anotherDate1.Sub(startDate).Hours()
	diffHr2 := anotherDate2.Sub(startDate).Hours()
	diffHr3 := targetDate1.Sub(startDate).Hours()
	diffHr4 := targetDate2.Sub(startDate).Hours()
	diffHr5 := anotherDate3.Sub(startDate).Hours()
	diffHr6 := targetDate3.Sub(startDate).Hours()

	t.Logf("%f", math.Mod(diffHr1-duration, thirtyDaysCycle))
	t.Logf("%f", math.Mod(diffHr2-duration, thirtyDaysCycle))
	t.Logf("%f", math.Mod(diffHr3-duration, thirtyDaysCycle))
	t.Logf("%f", math.Mod(diffHr4-duration, thirtyDaysCycle))
	t.Logf("%f", math.Mod(diffHr5-duration, thirtyDaysCycle))
	t.Logf("%f", math.Mod(diffHr6-duration, thirtyDaysCycle))
}

func TestAddDate(t *testing.T) {
	m := time.Now()
	t.Logf("%s", m.AddDate(0, -1, 0))
	t.Logf("%s", m.AddDate(0, 1, 0))

	d1 := time.Date(2017, 3, 29, 0, 0, 0, 0, time.Local)
	t.Logf("%s", d1)
	t.Logf("%s", d1.AddDate(0, -1, 0))

	d2 := time.Date(2017, 4, 29, 0, 0, 0, 0, time.Local)
	t.Logf("%s", d2)
	t.Logf("%s", d2.AddDate(0, -1, 0))

	d3 := time.Date(2017, 3, 28, 0, 0, 0, 0, time.Local)
	t.Logf("%s", d3)
	t.Logf("%s", d3.AddDate(0, -1, 0))

	d4 := time.Date(2017, 3, 30, 0, 0, 0, 0, time.Local)
	t.Logf("%s", d4)
	t.Logf("%s", d4.AddDate(0, -1, 0))

	d5 := time.Date(2017, 3, 31, 0, 0, 0, 0, time.Local)
	t.Logf("%s", d5)
	t.Logf("%s", d5.AddDate(0, -1, 0))

	d6 := time.Date(2017, 5, 31, 0, 0, 0, 0, time.Local)
	t.Logf("%s", d6)
	t.Logf("%s", d6.AddDate(0, -1, 0))

	d7 := time.Date(2017, 5, 31, 0, 0, 0, 0, time.UTC)
	t.Logf("%s", d7)
	t.Logf("%s", d7.AddDate(0, 1, 0))
}

func TestBetween(t *testing.T) {
	n := time.Date(2017, 4, 1, 18, 40, 0, 0, time.Local)
	noon := time.Date(n.Year(), n.Month(), n.Day(), 12, 0, 0, 0, time.Local)
	evening := time.Date(n.Year(), n.Month(), n.Day(), 18, 0, 0, 0, time.Local)

	switch {
	case n.Before(noon):
		t.Log("good morning")
	case n.After(noon) && n.Before(evening):
		t.Log("good afternoon")
	case n.After(evening):
		t.Log("good evening")
	}
	t.Logf("%s", n)
}
