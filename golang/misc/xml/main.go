package main

import (
	"encoding/xml"
	"log"
)

type Name struct {
	First string
	Last  string
}

type Person struct {
	Name Name
	Age  int
}

func main() {
	p1 := Person{
		Name: Name{
			First: "Akira",
			Last:  "Chiku",
		},
		Age: 30,
	}
	buf, _ := xml.MarshalIndent(p1, "", "  ")

	log.Printf("%v", p1)
	log.Printf("\n%s", string(buf))

	p2 := Person{}
	xmldoc := []byte(`
	<Person>
	  <Name>
		<First>Taro</First>
		<Last>Suzuki</Last>
	  </Name>
	  <Age>30</Age>
	</Person>
	`)
	xml.Unmarshal(xmldoc, &p2)
	log.Println(p2)
}
