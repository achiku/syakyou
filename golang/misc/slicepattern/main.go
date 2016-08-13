package main

import (
	"fmt"
	"strings"
)

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
