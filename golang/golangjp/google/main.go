package main

// http://talks.golang.org/2012/concurrency.slide#43
import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

var (
	Web   = fakeSearch("Web")
	Image = fakeSearch("Image")
	Video = fakeSearch("Video")
)

var (
	Web1   = fakeSearch("Web")
	Web2   = fakeSearch("Web")
	Image1 = fakeSearch("Image")
	Image2 = fakeSearch("Image")
	Video1 = fakeSearch("Video")
	Video2 = fakeSearch("Video")
)

func Google1(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return results
}

func Google2(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return results
}

func Google3(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return results
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func Google4(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return results
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google1("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Printf("Google1 Elapsed: %s\n", elapsed)

	rand.Seed(time.Now().UnixNano())
	start = time.Now()
	results = Google2("golang")
	elapsed = time.Since(start)
	fmt.Println(results)
	fmt.Printf("Google2 Elapsed: %s\n", elapsed)

	rand.Seed(time.Now().UnixNano())
	start = time.Now()
	results = Google3("golang")
	elapsed = time.Since(start)
	fmt.Println(results)
	fmt.Printf("Google3 Elapsed: %s\n", elapsed)

	rand.Seed(time.Now().UnixNano())
	start = time.Now()
	results = Google4("golang")
	elapsed = time.Since(start)
	fmt.Println(results)
	fmt.Printf("Google4 Elapsed: %s\n", elapsed)
}
