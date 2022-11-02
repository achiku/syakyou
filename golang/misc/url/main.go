package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u := "https://example.com?key1=val1#1&key2=val...2&url=https://akirachiku.com?mykey=val"
	pu, err := url.Parse(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pu)
	fmt.Println(url.QueryEscape(pu.String()))

	param := url.Values{}
	param.Set("key1", "val1")
	param.Set("key2", "val...2")
	param.Set("url", "https://akirachiku.com?mykey=val")
	fmt.Println(param.Encode())
}
