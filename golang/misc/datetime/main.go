package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	t := time.Now()
	t2 := t.AddDate(0, 0, 7)
	fmt.Println(t.Format("200601021504"))
	fmt.Println(t2.Format("200601021504"))

	tStr := "20160522203020" + " +0900 JST"
	t3, err := time.Parse("20060102150405 -0700 MST", tStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t3)
}
