package main

import "fmt"

// Handler interface
type Handler interface {
	ServeHTTP(string, string)
}

// HandlerStruct named handler
type HandlerStruct struct {
	Name string
}

func (h *HandlerStruct) ServeHTTP(w string, r string) {
	fmt.Printf("name: %s -> response: %s, request: %s\n", h.Name, w, r)
}

// HandlerFunc type
type HandlerFunc func(string, string)

func (f HandlerFunc) ServeHTTP(w string, r string) {
	f(w, r)
}

// Exe execute Handler interface
func Exe(handler Handler) {
	handler.ServeHTTP("test response", "test request")
}

func main() {
	var fn1 Handler
	fn1 = HandlerFunc(func(w string, r string) {
		fmt.Printf("response: %s, request: %s\n", w, r)
	})
	var fn2 Handler
	fn2 = &HandlerStruct{Name: "myHandler"}

	funcs := []Handler{fn1, fn2}
	for _, f := range funcs {
		Exe(f)
	}
}
