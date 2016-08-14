package main

import (
	"fmt"
	"sort"
	"strings"
)

// MyError my error
type MyError struct {
	orgError string
	errList  []string
}

// NewMyError new error
func NewMyError(s string) *MyError {
	return &MyError{
		orgError: s,
		errList:  strings.Split(s, "|"),
	}
}

// Error returns string
func (m *MyError) Error() string {
	return fmt.Sprintf("%s", m.errList)
}

// Has check if it has specified error
func (m *MyError) Has(s string) bool {
	for _, e := range m.errList {
		if e == s {
			return true
		}
	}
	return false
}

// Equals check if it equals specific combination of error
func (m *MyError) Equals(sl []string) bool {
	if len(m.errList) != len(sl) {
		return false
	}
	sort.Strings(m.errList)
	sort.Strings(sl)
	for i, v := range m.errList {
		if sl[i] != v {
			return false
		}
	}
	return true
}

type errList []string

func newErrList(s string) errList {
	l := strings.Split(s, "|")
	return errList(l)
}

func (el *errList) has(v string) bool {
	for _, e := range []string(*el) {
		if e == v {
			return true
		}
	}
	return false
}

func keys(m map[string]string) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	fmt.Println("vim-go")
}
