package main

import "net/http"

type Todo struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
type Todos []Todo

func main() {
	client := http.Client{}
}
