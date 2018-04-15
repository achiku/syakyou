package main

import "testing"

func TestTwice(t *testing.T) {
	a := twice(1, 11)
	t.Logf("%v", a)
}

func TestTwiceWithSlice(t *testing.T) {
	data := []int{1, 2, 3, 4}
	a := twice(data...)
	t.Logf("%v", a)
}

func TestTwiceWithSliceFunc(t *testing.T) {
	f := func() []int {
		return []int{1, 2, 3, 4}
	}
	a := twice(f()...)
	t.Logf("%v", a)
}

func TestTwiceWithSliceFuncAndSingleValue(t *testing.T) {
	f := func() []int {
		return []int{1, 2, 3, 4}
	}
	a := twice(append(f(), 1)...)
	t.Logf("%v", a)
}
