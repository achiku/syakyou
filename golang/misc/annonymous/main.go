package main

import "fmt"

type sts []struct {
	ID   string
	Name string
}

type st struct {
	ID   string
	Name string
}

func main() {
	var ss sts
	l := map[string]string{"achiku": "akira.chiku", "moqada": "masahiko.okada"}
	for k, v := range l {
		s := st{ID: k, Name: v}
		ss = append(ss, s)
	}
	fmt.Printf("%+v", ss)
}
