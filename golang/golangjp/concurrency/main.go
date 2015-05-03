package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Boring(msg string, c chan string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	times := 3
	c := make(chan string)
	go Boring("boring!", c)
	for i := 0; i < times; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}
