package main

import "log"

type customType string

var (
	typeA customType = "0A"
	typeB customType = "0B"
)

func main() {
	l := []string{"0A", "0B", "0X"}
	for _, i := range l {
		log.Println(i)
		switch customType(i) {
		case typeA:
			log.Println("this is typeA")
		case typeB:
			log.Println("this is typeB")
		default:
			log.Println("unknown type")
		}
	}
}
