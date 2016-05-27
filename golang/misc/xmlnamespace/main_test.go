package main

import (
	"encoding/xml"
	"testing"
)

func TestMarshalPerson(t *testing.T) {
	p := Person{
		ID: 10,
		Name: &Name{
			First: "Akira",
			Last:  "Chiku",
		},
		Age: 31,
	}
	buf, _ := xml.MarshalIndent(p, "", "  ")
	t.Logf("\n%s", string(buf))

	np := PersonNS{
		NS: "http://akirachiku.com/soap/",
		ID: 10,
		Name: &Name{
			First: "Akira",
			Last:  "Chiku",
		},
		Age: 31,
	}
	buf, _ = xml.MarshalIndent(np, "", "  ")
	t.Logf("\n%s", string(buf))
}
