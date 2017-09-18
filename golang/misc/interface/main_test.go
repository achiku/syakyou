package main

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type req struct {
	ID     interface{} `json:"id"`
	Amount int64       `json:"amount"`
}

const strJSON = `
{
  "id": "10222",
  "amount": 25000 
}
`

const intJSON = `
{
  "id": 10222,
  "amount": 25000 
}
`

const zeroIntJSON = `
{
  "id": 0,
  "amount": 25000 
}
`

const zeroStrJSON = `
{
  "id": "0",
  "amount": 25000 
}
`

func TestDecodeInterfaceJSON(t *testing.T) {
	cases := []string{strJSON, intJSON, zeroIntJSON, zeroStrJSON}
	for _, c := range cases {
		enc := json.NewDecoder(strings.NewReader(c))
		var r req
		if err := enc.Decode(&r); err != nil {
			t.Fatal(err)
		}
		var i int64
		switch r.ID.(type) {
		case float64:
			t.Logf("float")
			fi := r.ID.(float64)
			i = int64(fi)
		case string:
			si := r.ID.(string)
			i, _ = strconv.ParseInt(si, 10, 64)
			t.Logf("string")
		default:
			fv := reflect.TypeOf(r.ID)
			t.Logf("default: %s", fv)
		}
		t.Logf("%+v", r)
		t.Logf("%d", i)
	}
}

func TestMisc(t *testing.T) {
	var hh interface{}
	hh = float64(0)
	ii, ok := hh.(float64)
	if !ok {
		t.Fatal("failed")
	}
	i := int64(ii)
	t.Logf("%f, %d, %v", ii, i, hh)
}
