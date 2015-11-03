package main

import (
	"fmt"
	"math/rand"
	"time"
)

func helloAPI(name string, c chan string) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	greeting := "Hello, " + name + "!"
	c <- greeting
}

func profileAPI(name string, c chan string) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	profile := name + "'s profile"
	c <- profile
}

func main() {
	rand.Seed(time.Now().UnixNano())
	helloChan := make(chan string)
	profChan := make(chan string)
	users := []string{"8maki", "moqada", "ide", "achiku"}

	for _, u := range users {
		go helloAPI(u, helloChan)
		go profileAPI(u, profChan)
	}

	numAPI := 2
	for i := 0; i < len(users)*numAPI; i++ {
		select {
		case hello := <-helloChan:
			fmt.Println(hello)
		case profile := <-profChan:
			fmt.Println(profile)
		}
	}
	fmt.Println("done!")
}
