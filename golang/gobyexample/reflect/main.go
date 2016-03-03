package main

import (
	"log"
	"reflect"
)

func main() {
	// type MyType int
	// type MyType string
	// type MyType []string
	type MyType map[string]int
	var x MyType
	v := reflect.ValueOf(x)
	log.Println(v)
	log.Println(v.Type())
	log.Println(v.Kind())
}
