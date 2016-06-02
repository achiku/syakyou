package main

import (
	"fmt"
	"log"
	"time"
)

type person struct {
	Name     string
	Birthday time.Time
}

type person2 struct {
	Name     string
	Birthday *time.Time
}

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

	p1 := person{
		Name:     "achiku",
		Birthday: time.Now(),
	}
	fmt.Printf("%+v\n", p1)

	dob := time.Now()
	p2 := person2{
		Name:     "achiku",
		Birthday: &dob,
	}
	fmt.Printf("%+v\n", p2)
}
