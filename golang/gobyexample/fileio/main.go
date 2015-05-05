package main

import (
	"fmt"
	"log"
	"os"
)

type MyFile struct {
	*os.File
	Name string
	Stat os.FileInfo
}

func MyNewFile(path string) (*MyFile, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	stat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	file := &MyFile{
		f,
		path,
		stat,
	}

	return file, err
}

func main() {
	file, err := MyNewFile("./tmp")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf(
		"name=%s, size=%d, mode=%s, modtime=%s\n",
		file.Name,
		file.Stat.Size(),
		file.Stat.Mode(),
		file.Stat.ModTime(),
	)
}
