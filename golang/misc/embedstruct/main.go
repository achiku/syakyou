package main

import "log"

type A struct {
	Foo string
}

func (a A) Say(s string) string {
	return a.Foo
}

func (a A) GetFoo() string {
	return a.Foo
}

type Aish interface {
	Say(string) string
	GetFoo() string
}

type B struct {
	*A
	Bar string
}

func (b B) Say(s string) string {
	return b.A.Foo
}

func (b B) GetFoo() string {
	return b.A.Foo
}

func printAish(a Aish) string {
	return a.Say("aaaa")
}

func main() {
	a := A{Foo: "foo"}
	b := B{A: &A{Foo: "b foo"}}

	log.Println(printAish(a))
	log.Println(printAish(b))
}
