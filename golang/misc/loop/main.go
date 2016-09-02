package main

import "log"

func main() {
	log.Println("continue")
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range l {
		if i%3 == 0 {
			continue
		} else {
			log.Println(i)
		}
		log.Println("--")
	}

	log.Println("break")
	for _, i := range l {
		if i%3 == 0 {
			break
		} else {
			log.Println(i)
		}
		log.Println("--")
	}

	log.Println("labeled break")
OUTER:
	for _, i := range l {
		for _, j := range l {
			if i%2 == 0 && j%3 == 0 {
				break OUTER
			} else {
				log.Printf("i=%d, j=%d", i, j)
			}
			log.Println("------")
		}
	}
}
