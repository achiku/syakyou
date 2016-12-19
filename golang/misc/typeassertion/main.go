package main

import "log"

func returnString() interface{} {
	return "string"
}

func returnInt() interface{} {
	return 100
}

func main() {
	str, ok := returnString().(string)
	if !ok {
		log.Printf("expected string but got some thing wrong")
		return
	}
	log.Printf("%s", str)

	i, ok := returnInt().(int)
	if !ok {
		log.Printf("expected int but got some thing wrong")
		return
	}
	log.Printf("%d", i)
}
