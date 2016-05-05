package main

import (
	"log"
	"os"
)

func main() {
	h := os.Getenv("HOME")
	log.Println(h)
	nh := os.Getenv("NOT_HOME")
	log.Println(nh)
	if nh == "" {
		log.Println("nh is empty string")
	}
}
