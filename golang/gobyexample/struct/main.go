package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

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
