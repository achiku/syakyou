package main

import "testing"

func TestIncr(t *testing.T) {
	l := []int{1, 2, 3, 3, 3, 1}
	m := make(map[int]int)
	for _, i := range l {
		m[i]++
	}
	t.Logf("%v", m)
}
