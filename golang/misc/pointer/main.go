package main

// https://dhdersch.github.io/golang/2016/01/23/golang-when-to-use-string-pointers.html
// http://stackoverflow.com/questions/26493923/go-string-pointer-to-string

import "log"

type p struct {
	Name string
	Flag bool
}

func main() {
	var s *string
	log.Println(s)
	a := "aaaa"
	s = &a
	log.Println(s)
	log.Println(*s)
	log.Printf("%s", *s)

	var pi *p
	var pj *p
	if true {
		log.Printf("%+v", pi)
		log.Printf("%+v", pj)
		pi = &p{Name: "name", Flag: true}
	}
	log.Printf("%s", pi.Name)
	// log.Printf("%s", pj.Name) // panic!
}
