package main

import "log"

const i = iota

func main() {
	names := []string{"8maki", "ideyuta", "moqada"}
	for _, n := range names {
		log.Println(n)
		log.Println(i)
	}
}
