package main

import (
	"errors"
	"log"
	"time"
)

type person struct {
	Name string
	DOB  time.Time
}

func newPerson(name string, dob time.Time) (person, error) {
	var p person
	if name == "" {
		return p, errors.New("name is empty")
	}
	return person{
		Name: name,
		DOB:  dob,
	}, nil
}

func fn() (int, error) {
	return 1,
		errors.New("some error")
}

func main() {
	p, err := newPerson("achiku", time.Date(1990, time.January, 1, 0, 0, 0, 0, time.Local))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", p)

	i, err := fn()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(i)
}
