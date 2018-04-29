package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"strings"
	"testing"
)

func TestReadDataByCSV(t *testing.T) {
	fh := strings.NewReader(data)
	reader := csv.NewReader(fh)
	recs, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	for _, rec := range recs {
		fmt.Printf("%s\n", rec)
	}
}

func TestReadDataByScanner(t *testing.T) {
	fh := strings.NewReader(data)
	sc := bufio.NewScanner(fh)
	p := newParser()
	for sc.Scan() {
		p.parse(sc.Text())
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}
	for _, a := range p.f {
		t.Logf("%+v, %+v, %+v", a.rec1, a.rec2, a.rec3)
	}
}
