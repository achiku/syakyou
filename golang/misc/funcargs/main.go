package main

func twice(l ...int) []int {
	var t []int
	for _, i := range l {
		t = append(t, i*2)
	}
	return t
}
