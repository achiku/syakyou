package main

import "fmt"

func main() {
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Print("Continue\n")
	for i, n := range l {
		if n%2 == 0 {
			continue
		}
		fmt.Printf("%d: %d\n", i, n)
	}

	fmt.Print("Break\n")
	for i, n := range l {
		if n == 7 {
			break
		}
		fmt.Printf("%d: %d\n", i, n)
	}
}
