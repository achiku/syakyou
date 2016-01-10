// http://blog.monochromegane.com/blog/2014/03/23/struct-implementaion-patterns-in-golang/
package main

import "fmt"

// Company struct
type company struct {
	Name string
}

var kanmu *company

// GetKanmu return kanmu singleton instance
func getKanmu() *company {
	if kanmu == nil {
		fmt.Println("kanmu initialized")
		kanmu = &company{Name: "kanmu"}
	}
	return kanmu
}

// AnimalBehavior interface
type AnimalBehavior interface {
	Sleep(hours int) string
	Move(distance int) string
}

// Animal struct
type Animal struct {
	Weight int
	Height int
}

func (a *Animal) tellHeight() string {
	return fmt.Sprintf("%d height", a.Height)
}

func (a *Animal) tellWeight() string {
	return fmt.Sprintf("%d wieght", a.Weight)
}

// Person struct
type Person struct {
	*Animal
	Name string
	Age  int
}

// initialize is private to this package
func (p *Person) initialize() {
	fmt.Printf("user %s is initialized.\n", p.Name)
}

// greeting is private to this package
func (p *Person) greeting() {
	fmt.Printf("Hello! I'm %s!\n", p.Name)
}

// Move returns string
func (p *Person) Move(distance int) string {
	return fmt.Sprintf("%s walks %d meter", p.Name, distance)
}

// Sleep returns string
func (p *Person) Sleep(hours int) string {
	return fmt.Sprintf("%s sleeps %d hours", p.Name, hours)
}

// NewPerson create new person
func NewPerson(name string, age int) *Person {
	p := &Person{
		Name: name,
		Age:  age,
	}
	p.initialize()
	return p
}

// DoSomething func type
type DoSomething func()

// StructA struct
type StructA struct {
}

// Operation strategy
func (a *StructA) Operation(strategy DoSomething) {
	strategy()
}

func main() {
	teamMembers := []*Person{
		NewPerson("moqada", 30),
		NewPerson("8maki", 30),
		NewPerson("ideyuta", 27),
	}

	for _, p := range teamMembers {
		k := getKanmu()
		fmt.Printf("company name: %s\n", k.Name)
		p.greeting()
		fmt.Println(p.Move(10))
	}

	var animal AnimalBehavior
	animal = &Person{
		Name: "achiku",
		Age:  30,
	}
	// animal.greeting() -> this can't be called
	fmt.Println(animal.Move(10))
	fmt.Println(animal.Sleep(30))

	structA := &StructA{}
	structA.Operation(func() {
		fmt.Println("Operation")
	})
}
