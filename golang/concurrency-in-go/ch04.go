package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func pipeline2() {
	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		stream := make(chan interface{})
		go func() {
			defer close(stream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case stream <- v:
					}
				}
			}
		}()
		return stream
	}
	take := func(done <-chan interface{}, stream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-stream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	for num := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%d\n", num)
	}
}

func pipeline1() {
	mul := func(vals []int, m int) []int {
		results := make([]int, len(vals))
		for i, v := range vals {
			results[i] = v * m
		}
		return results
	}
	add := func(vals []int, m int) []int {
		results := make([]int, len(vals))
		for i, v := range vals {
			results[i] = v + m
		}
		return results
	}
	ints := []int{1, 2, 3, 4}
	for _, v := range add(mul(ints, 2), 1) {
		fmt.Println(v)
	}
}

func errorHnadling() {
	type result struct {
		er   error
		resp *http.Response
	}
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan result {
		results := make(chan result)
		go func() {
			defer close(results)
			for _, url := range urls {
				resp, err := http.Get(url)
				res := result{er: err, resp: resp}
				select {
				case <-done:
					return
				case results <- res:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://badhost"}
	for res := range checkStatus(done, urls...) {
		if res.er != nil {
			fmt.Printf("error: %s\n", res.er)
			continue
		}
		fmt.Printf("Response: %v\n", res.resp.Status)
	}
}

func confinement5() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWrok goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("done")
}

func confinement4() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)
	fmt.Println("done")
}

func confinement3() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	wg.Wait()
}

func confinement2() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for res := range results {
			fmt.Printf("Received: %d\n", res)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)
}

func confinement() {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}
