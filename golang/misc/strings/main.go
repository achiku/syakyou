package main

import "strings"

func getLastFour(s string) string {
	l := strings.Split(s, "")
	if len(l) < 4 {
		return s
	}
	return strings.Join(l[len(l)-4:], "")
}
