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

func tp(t time.Time) *time.Time {
	return &t
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

	p2 := person2{
		Name:     "achiku",
		Birthday: tp(time.Now()),
	}
	fmt.Printf("%+v\n", p2)

	p3 := person2{
		Name:     "achiku",
		Birthday: tp(time.Date(1985, 8, 18, 8, 8, 8, 8, time.Local)),
	}
	fmt.Printf("%+v\n", p3)
}
