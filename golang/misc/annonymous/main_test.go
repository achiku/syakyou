package main

import "testing"

func TestErrors(t *testing.T) {
	e := Error{
		Code: "suspended",
		Errors: Errors{
			{Code: "22112"},
		},
	}
	t.Logf("%v", e)
}
