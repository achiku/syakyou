package main

import (
	"log"
	"time"
)

func main() {
	t := time.Now()
	t2 := t.AddDate(0, 0, 7)
	log.Println(t.Format("200601021504"))
	log.Println(t2.Format("200601021504"))
}
