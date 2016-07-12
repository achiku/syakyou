package main

// https://dhdersch.github.io/golang/2016/01/23/golang-when-to-use-string-pointers.html
// http://stackoverflow.com/questions/26493923/go-string-pointer-to-string

import "log"

func main() {
	var s *string
	log.Println(s)
	a := "aaaa"
	s = &a
	log.Println(s)
	log.Println(*s)
	log.Printf("%s", *s)
}
