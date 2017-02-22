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

// String string
func (s SecuerStruct) String() string {
	return fmt.Sprintf("{Name: %s, Password: ****}", s.Name)
}
