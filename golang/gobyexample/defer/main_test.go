package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var testFileContent = "this is a test file\n"

func TestBadCopyFile(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	if err := ioutil.WriteFile(tmpFile.Name(), []byte(testFileContent), 0644); err != nil {
		t.Error("should not raise error: %v", err)
	}

	written, err := BadCopyFile("./testfile.txt", tmpFile.Name())
	if err != nil {
		t.Error("BadCopyFile failed: %s, %s", err, written)
	}
}

func TestGoodCopyFile(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	if err := ioutil.WriteFile(tmpFile.Name(), []byte(testFileContent), 0644); err != nil {
		t.Error("should not raise error: %v", err)
	}

	written, err := GoodCopyFile("./testfile.txt", tmpFile.Name())
	if err != nil {
		t.Error("GoodCopyFile failed: %s, %s", err, written)
	}
}

func TestPrintNum01(t *testing.T) {
	fmt.Println("i will not be evaluated until the last fmt.Print() is executed")
	PrintNum01()
	fmt.Println()
}

func TestPrintNum02(t *testing.T) {
	PrintNum02()
	fmt.Println()
}
