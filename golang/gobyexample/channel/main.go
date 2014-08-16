package main

import (
	"fmt"
	"time"
)

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	/* channel buffering */
	messages := make(chan string)

	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)

	/* channel buffering */
	buf_messages := make(chan string, 2)
	buf_messages <- "buffered"
	buf_messages <- "channel"
	fmt.Println(<-buf_messages)
	fmt.Println(<-buf_messages)

	/* channel synchronization */
	done := make(chan bool, 1)
	go worker(done)

	/* If you removed the <- done line from this program,
	the program would exit before the worker even started */
	<-done

	/* channel synchronization */
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
