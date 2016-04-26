package main

import (
	"errors"
	"log"
)

func fn() (int, error) {
	return 1,
		errors.New("y")
}

func main() {
	i, err := fn()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(i)
}
