package main

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestCaller(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("error")
	}
	f, err := filepath.Abs(filename)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", f)
}

func TestPWD(t *testing.T) {
	d, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", d)
	f, err := filepath.Abs(d)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", f)
}
