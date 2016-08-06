//go:generate go-bindata -o template.go template
package main

import (
	"bytes"
	"go/format"
	"log"
	"text/template"

	"github.com/pkg/errors"
)

const structTmpl = `
// {{ .Name }} represents {{ .Schema }}.{{ .TableName }}
type {{ .Name }} struct {
{{- range .Fields }}
	{{ .Name }} {{ .Type }} // {{ .ColName }}
{{- end }}
}`

// StructField go struct field
type StructField struct {
	Name    string
	Type    string
	Tag     string
	NilVal  string
	ColName string
}

// Struct go struct
type Struct struct {
	Name      string
	TableName string
	Schema    string
	Comment   string
	Fields    []*StructField
}

func generate(st *Struct) ([]byte, error) {
	var src []byte
	tpl, err := template.New("struct").Parse(structTmpl)
	if err != nil {
		return src, errors.Wrap(err, "failed to parse template")
	}
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, st); err != nil {
		return src, errors.Wrap(err, "failed to execute template")
	}
	src, err = format.Source(buf.Bytes())
	if err != nil {
		return src, errors.Wrap(err, "failed to format source code")
	}
	return src, nil
}

func generateFromFile(st *Struct) ([]byte, error) {
	var src []byte
	tpl, err := template.ParseFiles("template/struct.tmpl")
	if err != nil {
		return src, errors.Wrap(err, "failed to parse template")
	}
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, st); err != nil {
		return src, errors.Wrap(err, "failed to execute template")
	}
	src, err = format.Source(buf.Bytes())
	if err != nil {
		return src, errors.Wrap(err, "failed to format source code")
	}
	return src, nil

}

func generateFromBinData(st *Struct) ([]byte, error) {
	var src []byte
	d, err := Asset("template/struct.tmpl")
	if err != nil {
		return src, errors.Wrap(err, "failed to get template from bindata")
	}
	tpl, err := template.New("struct").Parse(string(d))
	if err != nil {
		return src, errors.Wrap(err, "failed to parse template")
	}
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, st); err != nil {
		return src, errors.Wrap(err, "failed to execute template")
	}
	src, err = format.Source(buf.Bytes())
	if err != nil {
		return src, errors.Wrap(err, "failed to format source code")
	}
	return src, nil
}

func main() {
	log.Println("hello")
}
