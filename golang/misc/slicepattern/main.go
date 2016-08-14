package main

import (
	"fmt"
	"strings"
)

// Brand brand
type Brand struct {
	Type string
	Name string
}

func find(num string) *Brand {
	switch {
	case num[:1] == "4":
		return &Brand{
			Type: "visa",
			Name: "Visa",
		}
	case num[:1] == "5":
		return &Brand{
			Type: "master",
			Name: "Master Card",
		}
	case num[:2] == "34" || num[:2] == "37":
		return &Brand{
			Type: "amex",
			Name: "American Express",
		}
	default:
		return &Brand{
			Type: "unknown",
			Name: "unknown",
		}
	}
}

func main() {
	cmd := "ls"
	options := []string{"-l", "-a", "-h"}
	commandLine := append([]string{cmd}, options...)
	fmt.Println(options)
	fmt.Println(commandLine)
	fmt.Println(strings.Join(options, " "))

	s := "1234567890"
	fmt.Println(s[3:])
	fmt.Println(s[:1])
	fmt.Println(s[:len(s)-1])
	fmt.Println(s[:len(s)-3])

	i := []int{0, 1, 2, 3}
	fmt.Println(i[len(i)-1])
}
