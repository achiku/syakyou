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

func TestOverwriteField(t *testing.T) {
	bv := b{
		SecuerStruct: SecuerStruct{
			Name:     "this is secure struct name",
			Password: "secure password",
		},
		Name: "this is b name",
	}
	t.Logf("%s", bv.Name)
	t.Logf("%s", bv.SecuerStruct.Name)
}
