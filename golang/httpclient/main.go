package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	req, err := http.NewRequest("GET", "http://localhost:8080/todo/", nil)
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

	for k, v := range resp.Header {
		fmt.Println(k, v)
	}

	bodyJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var todos Todos
	err = json.Unmarshal([]byte(bodyJson), &todos)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(todos)
	for _, t := range todos {
		fmt.Println(t)
	}
}
