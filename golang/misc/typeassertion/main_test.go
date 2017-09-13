package main

import (
	"log"
	"reflect"
	"testing"
)

func TestMiscMain(t *testing.T) {
	str, ok := returnString().(string)
	if !ok {
		log.Printf("expected string but got some thing wrong")
		return
	}
	log.Printf("%s", str)

	i, ok := returnInt().(int)
	if !ok {
		log.Printf("expected int but got some thing wrong")
		return
	}
	log.Printf("%d", i)

	a, ok := returnTestType().(testType)
	if !ok {
		log.Printf("expected testType but got some thing wrong")
		return
	}
	log.Printf("%s", a)
}

func TestStruct(t *testing.T) {
	ts := testStruct{
		ID:   100,
		Name: "test struct",
	}
	ats := anotherTestStruct(ts)
	t.Logf("%+v", ats)
	t.Logf("%s", reflect.TypeOf(ats))
	t.Logf("%s", reflect.TypeOf(ts))
}

type myType int

const (
	type1 myType = iota
	type2
)

func TestTypes(t *testing.T) {
	t.Logf("%s", reflect.TypeOf(type1))
	t.Logf("%s", reflect.TypeOf(type2))
}
