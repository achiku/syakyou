package main

import (
	"fmt"
	"log"
)

// StatusCode status
type StatusCode int

// Message message
type Message string

// HTTPStatus http status code
type HTTPStatus map[StatusCode]Message

var (
	codes = [...]HTTPStatus{
		HTTPStatus{200: "OK"},
		HTTPStatus{500: "Internal Server Error"},
	}
)

// Hex hex
type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("0x%x", int(h))
}

// Person person
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Name name
type Name struct {
	FirstName string
	LastName  string
}

func (n *Name) String() string {
	return fmt.Sprintf("%s %s", n.FirstName, n.LastName)
}

// Person2 person2
type Person2 struct {
	*Name
	Age int
}

// Person3 person3
type Person3 interface {
	Name() string
	Title() string
}

type person struct {
	firstName string
	lastName  string
}

// Gender gender
type Gender int

// Gender types
const (
	Female = iota
	Male
)

type female struct {
	*person
}

func (f *female) Title() string {
	return "Ms. "
}

type male struct {
	*person
}

func (m *male) Title() string {
	return "Mr. "
}

// NewPerson create person
func NewPerson(gender Gender, firstName, lastName string) Person3 {
	p := &person{firstName, lastName}
	if gender == Female {
		return &female{p}
	}
	return &male{p}
}

func (p *person) Name() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p *Person) String() string {
	return fmt.Sprintf("%s %s (%d)", p.FirstName, p.LastName, p.Age)
}

func main() {
	members := []Person{
		Person{LastName: "Okada", FirstName: "Masahiko", Age: 30},
		Person{LastName: "Yamaki", FirstName: "Wataru", Age: 30},
	}
	for _, u := range members {
		log.Println(u.String())
	}

	var h Hex
	h = 10
	log.Println(h.String())

	var stringer fmt.Stringer
	stringer = Hex(100)
	log.Println(stringer.String())

	n := &Name{FirstName: "Akira", LastName: "Chiku"}
	p := &Person2{Name: n, Age: 30}
	log.Println(p.String())

	var pstringer fmt.Stringer
	pstringer = Person2{Name: n, Age: 30}
	log.Println(pstringer.String())
}
