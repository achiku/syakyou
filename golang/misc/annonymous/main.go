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

// Errors errors
type Errors []struct {
	Detail     string `json:"detail,omitempty"`
	UserDetail string `json:"user_detail,omitempty"`
	Field      string `json:"field,omitempty"`
	Code       string `json:"code,omitempty"`
}

// Error struct for error resource
type Error struct {
	Code   string `json:"code,omitempty"`
	Detail string `json:"detail,omitempty"`
	Errors []struct {
		Detail     string `json:"detail,omitempty"`
		UserDetail string `json:"user_detail,omitempty"`
		Field      string `json:"field,omitempty"`
		Code       string `json:"code,omitempty"`
	} `json:"errors,omitempty"`
	Status     int64  `json:"status"`
	Title      string `json:"title,omitempty"`
	Type       string `json:"type"`
	UserDetail string `json:"user_detail,omitempty"`
	UserTitle  string `json:"user_title,omitempty"`
}

func main() {
	var ss sts
	l := map[string]string{"achiku": "akira.chiku", "moqada": "masahiko.okada"}
	for k, v := range l {
		s := stt{ID: k, Name: v}
		ss = append(ss, s)
	}
	a := st{
		ID:   "achiku",
		Name: "akira.chiku",
	}
	fmt.Printf("%+v", ss)
	fmt.Printf("%+v", a)
}
