package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	// 2006-01-02 15:04:05
	r1 := `{"name": "achiku", "created_at": "2014-09-03 10:20:01 UTC"}`

	var req1 req
	decoder := json.NewDecoder(strings.NewReader(r1))
	if err := decoder.Decode(&req1); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", req1)
}
