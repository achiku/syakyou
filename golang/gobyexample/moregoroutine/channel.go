package main

import "log"

func f(ch chan bool) {
	ch <- true
}

func main() {
	ch := make(chan bool)
	go f(ch)
	log.Println(<-ch)

	fin := make(chan bool)
	go func() {
		log.Println("working...")
		fin <- false
	}()
	<-fin
}
