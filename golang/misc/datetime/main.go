package main

import (
	"fmt"
	"log"
	"strconv"
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

func firstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func lastDayOfMonth(t time.Time) time.Time {
	startMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	endMonth := startMonth.AddDate(0, 1, -1)
	endMonth = time.Date(
		endMonth.Year(), endMonth.Month(), endMonth.Day(), 23, 59, 59, 0, endMonth.Location())
	return endMonth
}

func main() {
	t := time.Now()
	fmt.Printf("mmyy: %s\n", t.Format("0106"))
	t2 := t.AddDate(0, 0, 7)
	fmt.Println(t.Format("15:04:05"))
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

	yymm := "2601"
	yymmT, err := time.Parse("0601", yymm)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Printf("yymm %s", yymmT)
	fmt.Printf("%s\n", lastDayOfMonth(yymmT))

	a := time.Date(2016, 8, 13, 0, 0, 0, 0, time.Local)
	fmt.Printf("%s\n", a)

	// ldmCases := struct {
	// 	t time.Time
	// 	expected time.Time
	// } {
	// 	{t}
	// }
	y, _ := strconv.ParseInt(a.Format("06"), 10, 64)
	fmt.Println(y)
	tStr = "2016-08-18 22:08:54"
	tm1, err := time.Parse("2006-01-02 15:04:05", tStr)
	if err != nil {
		log.Fatal(err)
	}
	tm2, err := time.ParseInLocation("2006-01-02 15:04:05", tStr, time.Local)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(tm1)
	log.Println(tm2)
	log.Println(firstDayOfMonth(t2))

	n := time.Now()
	hho := time.Now().Add(10 * time.Minute)
	fmt.Printf("%s : %s : %s \n", n, n.Add(10*time.Minute), hho)
}
