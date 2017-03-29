package main

import (
	"fmt"
	"time"
)

// https://github.com/golang/go/issues/10401

// AddDate returns the time corresponding to adding the
// given number of years, months, and days to t.
// For example, AddDate(-1, 2, 3) applied to January 1, 2011
// returns March 4, 2010.
//
// AddDate normalizes its result in the same way that Date does,
// so, for example, adding one month to October 31 yields
// December 1, the normalized form for November 31.

func main() {
	d1 := time.Date(2017, 3, 29, 0, 0, 0, 0, time.UTC)
	fmt.Printf("%s\n", d1)
	fmt.Printf("%s\n", d1.AddDate(0, -1, 0))

	d2 := time.Date(2017, 3, 28, 0, 0, 0, 0, time.UTC)
	fmt.Printf("%s\n", d2)
	fmt.Printf("%s\n", d2.AddDate(0, -1, 0))
}
