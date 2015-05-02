package main

import "fmt"

type Values map[string]float64

func main() {

	m := make(map[string]int)
	v := make(Values)

	v["key01"] = 90.10
	v["key02"] = 110.12
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	fmt.Println("map:", v)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
