package main

// http://blog.golang.org/defer-panic-and-recover

import (
	"fmt"
	"io"
	"os"
)

/*
This works, but there is a bug. If the call to os.Create fails,
the function will return without closing the source file.
This can be easily remedied by putting a call to src.Close before the second return statement,
but if the function were more complex the problem might not be so easily noticed and resolved.
*/
func BadCopyFile(dstName string, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return 0, err
	}

	dst, err := os.Create(dstName)
	if err != nil {
		return 0, err
	}

	written, err = io.Copy(dst, src)
	if err != nil {
		return 0, err
	}
	dst.Close()
	src.Close()
	return written, nil
}

/*
Defer statements allow us to think about closing each file right after opening it,
guaranteeing that, regardless of the number of return statements in the function, the files will be closed.
*/
func GoodCopyFile(dstName string, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return 0, err
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func PrintNum01() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func PrintNum02() {
	i := 0
	defer fmt.Println("defer01: ", i)
	i++
	defer fmt.Println("defer02: ", i)
	fmt.Println("non defer:", i)
	return
}
