package main

import (
	"encoding/xml"
	"fmt"
)

// Name struct
type Name struct {
	XMLName xml.Name `xml:"http://example.com/ns1 name"`
	First   string   `xml:"first,omitempty"`
	Last    string   `xml:"last,omitempty"`
}

// Person struct
type Person struct {
	XMLName xml.Name `xml:"http://example.com/ns1 person"`
	ID      int      `xml:"id,omitempty"`
	Name    *Name
	Age     int `xml:"age,omitempty"`
}

// PersonNS struct
type PersonNS struct {
	XMLName xml.Name `xml:"ns:Person"`
	NS      string   `xml:"ns,attr"`
	ID      int      `xml:"ns:Id,omitempty"`
	Name    *Name
	Age     int `xml:"ns:Age,omitempty"`
}

// PersonNS2 struct
type PersonNS2 struct {
	XMLName xml.Name `xml:"ns:Person"`
	ID      int      `xml:"ns:Id,omitempty"`
	Name    *Name
	Age     int `xml:"ns:Age,omitempty"`
}

func main() {
	fmt.Printf("Hello, world\n")
}
