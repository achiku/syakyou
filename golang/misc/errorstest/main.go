package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("Hello, world\n")
	fmt.Printf("%v\n", errors.New("aa"))
	e := errors.New("")
	if e == nil {
		fmt.Printf("e is nil\n")
	}
	switch t := e.(type) {
	default:
		fmt.Printf("unexpected type %T", t)
	case error:
		fmt.Printf("error type %T", t)
	}
}
