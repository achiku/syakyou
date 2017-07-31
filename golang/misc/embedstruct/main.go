package main

import "fmt"

type a struct {
	SecuerStruct
	AName string
}

// SecuerStruct secure
type SecuerStruct struct {
	Name     string
	Password string
}

type b struct {
	SecuerStruct
	Name string
}

// String string
func (s SecuerStruct) String() string {
	return fmt.Sprintf("{Name: %s, Password: ****}", s.Name)
}

type name struct {
	First string
	Last  string
}

type child struct {
	Name name
}

type parent struct {
	Name  name
	Child *child
}
