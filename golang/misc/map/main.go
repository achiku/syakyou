package main

import "fmt"

func main() {
	mp := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	for k, v := range mp {
		fmt.Printf("%s->%d\n", k, v)
	}

	for _, k := range []string{"key1", "key2", "key3", "key4"} {
		v, ok := mp[k]
		if !ok {
			fmt.Printf("value for key (%s) not found\n", k)
			continue
		}
		fmt.Printf("value (%d) for key (%s)\n", v, k)
	}
}
