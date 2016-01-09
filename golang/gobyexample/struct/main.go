package main

import (
	"fmt"
	"strconv"
)

// Person struct
type Person struct {
	name string
	age  int
}

func (person *Person) hello() string {
	return "hello, I'm " + person.name + ", and " + strconv.Itoa(person.age) + " years old."
}

func main() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Alice", age: 30})

	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age: 40})

	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)
	fmt.Println(s.hello())

	/* this is just copying struct data to a variable a */
	a := s
	fmt.Println(a.age)
	fmt.Println(s.age)

	a.age = 99
	fmt.Println(a.age)
	fmt.Println(s.age)

	/* change value inside a struct using pointer */
	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
