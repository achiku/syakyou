package main

import "fmt"

// SecuerStruct secure
type SecuerStruct struct {
	Name     string
	Password string
}

// String string
func (s SecuerStruct) String() string {
	return fmt.Sprintf("{Name: %s, Password: ****}", s.Name)
}
