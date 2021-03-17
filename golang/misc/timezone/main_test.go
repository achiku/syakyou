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
	// この下は0になる
	t.Logf("%s", tm.Sub(tmJST))
}

func TestCompare(t *testing.T) {
	tm := time.Date(2017, 8, 18, 0, 0, 0, 0, time.Local)
	tn := time.Date(2017, 8, 12, 0, 0, 0, 0, time.Local)

	t.Logf("%t", tm.Before(tn))
	t.Logf("%t", tm.After(tn))
}

func TestTimePointer(t *testing.T) {
	t.Logf("hello")
	type s struct {
		tm *time.Time
	}

	now := time.Now()
	st := s{
		tm: &now,
	}
	t.Logf("%+v", st)
	t.Logf("%+v", st.tm)
}
