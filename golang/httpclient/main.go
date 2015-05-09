package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type Todo struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
type Todos []Todo

func main() {
	client := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	dumps, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dumps)
}
