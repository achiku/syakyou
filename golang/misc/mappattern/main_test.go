package main

import "testing"

func TestMap(t *testing.T) {
	cases := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
	}

	for k, v := range cases {
		t.Logf("%s => %s", k, v)
	}

	keys := keys(cases)
	keys = append(keys, "notkey")
	for _, k := range keys {
		if val, ok := cases[k]; ok {
			t.Logf("%s, %s", k, val)
		} else {
			t.Logf("key: %s not found", k)
		}
	}

}

func TestSplit(t *testing.T) {
	cases := []struct {
		s   string
		key string
		has bool
	}{
		{s: "", key: "E12345", has: false},
		{s: "E12345", key: "E12345", has: true},
		{s: "E12345", key: "E00000", has: false},
		{s: "E12345|E122222|E193837", key: "E193837", has: true},
		{s: "E12345|E122222|E193837", key: "E000000", has: false},
	}
	for _, c := range cases {
		e := newErrList(c.s)
		if e.has(c.key) != c.has {
			t.Errorf("%v, %s", e, c.key)
		}
	}
}
