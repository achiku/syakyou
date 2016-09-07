package main

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	l := []string{"A", "B", "C", "D"}
	for _, s := range l {
		switch s {
		case "A":
			fmt.Printf("A found\n")
		case "B":
			fmt.Printf("B found\n")
		case "C":
			fmt.Printf("C found\n")
		default:
			fmt.Printf("%s not found\n", s)
		}
	}
}
