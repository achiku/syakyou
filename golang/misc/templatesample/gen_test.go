package main

import "testing"

var fs = []*StructField{
	&StructField{Name: "col1", Type: "string", Tag: "", NilVal: "", ColName: "col_1"},
	&StructField{Name: "col2", Type: "int", Tag: "", NilVal: "0", ColName: "col_2"},
	&StructField{Name: "col3", Type: "bool", Tag: "", NilVal: "false", ColName: "col_3"},
}
var st = &Struct{
	Name:      "St1",
	TableName: "st_1",
	Schema:    "public",
	Fields:    fs,
}

func TestGenerate(t *testing.T) {
	src, err := generate(st)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", src)
}

func TestGenerateFromFile(t *testing.T) {
	src, err := generateFromFile(st)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", src)
}

func TestGenerateFromBinData(t *testing.T) {
	src, err := generateFromBinData(st)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", src)
}
