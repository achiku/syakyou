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
