package main

func getSortedStringKeys(in interface{}) []string {
	res, ok := in.(map[string]interface{})
	if !ok {
		panic("ug")
	}
	var keys []string
	for key := range res {
		keys = append(keys, key)
	}
	return keys
}
