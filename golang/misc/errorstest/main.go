package main

import (
	"fmt"
	"log"
)

type myErrorOne struct {
	Msg string
}

func (e myErrorOne) Error() string {
	return fmt.Sprintf("error one: msg=%s", e.Msg)
}

type myErrorTwo struct {
	Msg string
}

func (e myErrorTwo) Error() string {
	return fmt.Sprintf("error two: msg=%s", e.Msg)
}

func returnError(i int) error {
	if i == 1 {
		return myErrorOne{Msg: "fuck"}
	} else if i == 2 {
		return myErrorTwo{Msg: "shit"}
	}
	return nil
}

func main() {
	if err := returnError(1); err != nil {
		switch err := err.(type) {
		case myErrorOne:
			log.Printf("my error one: %s", err.Error())
		case myErrorTwo:
			log.Printf("my error one: %s", err.Error())
		}
	}
}
