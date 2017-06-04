package main

import "fmt"

type sts []struct {
	ID   string
	Name string
}

type stt struct {
	ID   string
	Name string
}

type st struct {
	ID      string
	Name    string
	Address []struct {
		State string
		City  string
	}
}

type address struct {
	State string
	City  string
}

func main() {
	var ss sts
	l := map[string]string{"achiku": "akira.chiku", "moqada": "masahiko.okada"}
	for k, v := range l {
		s := stt{ID: k, Name: v}
		ss = append(ss, s)
	}
	ad := []address{
		address{State: "Texas", City: "Austin"},
		address{State: "Texas", City: "Dallas"},
	}
	a := st{
		ID:      "achiku",
		Name:    "akira.chiku",
		Address: ad,
	}
	fmt.Printf("%+v", ss)
	fmt.Printf("%+v", a)
}
