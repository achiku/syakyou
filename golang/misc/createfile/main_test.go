package main

import (
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	path := "./test.data"
	src := []byte("test test test")

	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if err := write(f, src); err != nil {
		t.Fatal(err)
	}
}

func TestWriteStdout(t *testing.T) {
	src := []byte("test test test")

	f := os.Stdout
	if err := write(f, src); err != nil {
		t.Fatal(err)
	}
}
