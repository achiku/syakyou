package main

// https://github.com/go-fsnotify/fsnotify/blob/master/example_test.go

import (
	"log"
	"os"
	"path/filepath"

	"github.com/howeyc/fsnotify"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Usage: go run " + filepath.Base(os.Args[0]) + ".go [file_path]")
		os.Exit(1)
	}

	filePath := os.Args[1]

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Event:
				log.Println("event: ", event)
			}
		}
	}()

	err = watcher.Watch(filePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	<-done
}
