package main

import "testing"

func TestEmbedAndString(t *testing.T) {
	s := SecuerStruct{
		Name:     "achiku",
		Password: "secret",
	}
	t.Log(s)
}

func TestEmbed(t *testing.T) {
	st := a{
		SecuerStruct: SecuerStruct{
			Name:     "aaaaa",
			Password: "pass",
		},
		AName: "hogehoge",
	}
	t.Logf("%+v", st)
	t.Logf("%+v", st.SecuerStruct)
}
