package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println(os.Getpid())
	log.Printf("HOME: %s", os.Getenv("HOME"))
	log.Printf("USER: %s", os.Getenv("USER"))
	log.Printf("HOGE: %s", os.Getenv("HOGE"))
	log.Printf("APP_ENV: %s", os.Getenv("APP_ENV"))

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
