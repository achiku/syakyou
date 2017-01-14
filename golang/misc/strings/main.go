package main

import (
	"errors"
	"fmt"
	"strings"
)

func getLastFour(s string) string {
	l := strings.Split(s, "")
	if len(l) < 4 {
		return s
	}
	return strings.Join(l[len(l)-4:], "")
}

func getLastFour2(s string) (string, error) {
	if len(s) < 4 {
		return "", errors.New("s is less than four chars")
	}
	return s[len(s)-4:], nil
}

func trim(n string) string {
	r := []rune(n)
	namePart := string(r[:23])
	cleanName := strings.TrimRight(strings.TrimRight(namePart, " "), "　")
	return cleanName
}

func createJSON(amount string) []byte {
	s := fmt.Sprintf(`{"amount": "%s"}`, amount)
	return []byte(s)
}

func createJSON2(amount string) []byte {
	s := fmt.Sprintf("{\"amount\": \"%s\"}", amount)
	return []byte(s)
}
