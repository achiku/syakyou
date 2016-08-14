package main

import (
	"fmt"
	"strings"
)

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
