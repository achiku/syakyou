package main

// http://talks.golang.org/2012/concurrency.slide#26
import (
	"fmt"
	"math/rand"
	"time"
)

func Boring2(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- fmt.Sprintf("%s", msg):
			case <-quit:
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

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

func FanIn2(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
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

	c3 := FanIn2(Boring("achiku"), Boring("moqada"))
	for i := 0; i < 5; i++ {
		fmt.Println(<-c3)
	}
	fmt.Println("You're boring: I'm leaving.")

	// quit channel
	quit := make(chan bool)
	ide := Boring2("boring ide", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-ide)
	}
	quit <- true
	fmt.Println("Quit.")

	c4 := Boring("8maki")
	timeout := time.After(2 * time.Second)
	for {
		select {
		case s := <-c4:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}
