package main

import "testing"

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	// data := []int{1, 2, 3, 4}
	a := twice(1, 11)
	t.Logf("%v", a)
}
