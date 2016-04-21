package main

import (
	"fmt"
	"log"
	"net/http"
)

func postTest(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("key: %s", r.FormValue("key"))
	log.Printf("id: %s", r.FormValue("id"))
	log.Printf("message: %s", r.FormValue("message"))
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprint(w, "0000")
	return
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/post", postTest)
	return mux
}

func newServer(port string) *http.Server {
	mux := newMux()
	server := &http.Server{
		Handler: mux,
		Addr:    "localhost:" + port,
	}
	return server
}

func main() {
	port := "8899"
	s := newServer(port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
