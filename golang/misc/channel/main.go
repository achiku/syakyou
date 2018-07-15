package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func sayHello(i int, c chan string) {
	log.Printf("%d: hello!", i)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	c <- fmt.Sprintf("%d: hello!", i)
}

func main() {
	cnt := 5
	hc := make(chan string, cnt)
	for i := 0; i < cnt; i++ {
		go sayHello(i, hc)
		go func(i int) {
			log.Printf("%d: anonymous hello", i)
		}(i)
	}
	timeout := time.After(5 * time.Second)
	for {
		select {
		case msg := <-hc:
			log.Printf("hc: %s", msg)
		case <-timeout:
			log.Print("timeout")
			return
		}
	}
}
