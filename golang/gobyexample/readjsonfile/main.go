package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Message struct {
	User string `json:"user"`
	List []int  `json:"list"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./testdata_list.json")
	check(err)
	fmt.Println(string(dat))

	var ml []Message
	json.Unmarshal(dat, &ml)

	for i, m := range ml {
		fmt.Println(i, m.User)
	}

	for i, m := range ml {
		fmt.Println(i, m.List)
	}
}
