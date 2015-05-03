package main

// http://talks.golang.org/2012/concurrency.slide#26
import (
	"fmt"
	"math/rand"
	"time"
)

func Boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	// Generator: function that returns a channel
	// Channels are first-class values, just like strings or integers.
	c1 := Boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c1)
	}
	fmt.Println("You're boring: I'm leaving.")

	// Our boring function returns a channel that lets us communicate with the boring service it provides.
	// We can have more instances of the service.
	joe := Boring("Joe")
	ann := Boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're boring: I'm leaving.")

	// Multiplexing
	// above piece of programs make Joe and Ann count in lockstep.
	// We can instead use a fan-in function to let whosoever is ready talk.
	c2 := FanIn(Boring("Bob"), Boring("Alice"))
	for i := 0; i < 5; i++ {
		fmt.Println(<-c2)
	}
	fmt.Println("You're boring: I'm leaving.")
}
