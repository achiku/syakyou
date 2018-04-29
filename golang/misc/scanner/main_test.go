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
	lineNum := 1
	var (
		tran *transaction
		f    file
	)
	for sc.Scan() {
		l := sc.Text()
		d := strings.Split(l, ",")
		rec := record{
			category:     d[0],
			content:      d[1],
			categoryData: d[2],
		}
		// if it's still one transaction
		// if category ends with 0, it's start of the new transaction
		// category must be ascending order in the file
		typ := rec.categoryType()
		switch typ {
		case "0":
			tran = &transaction{
				rec1: &rec,
			}
			f = append(f, tran)
		case "1":
			tran.rec2 = &rec
		case "2":
			tran.rec3 = &rec
		}
		// add to rec2, rec3 of the transaction
		fmt.Printf("%d: %+v\n", lineNum, tran)
		lineNum++
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}

	for _, a := range f {
		t.Logf("%+v, %+v, %+v", a.rec1, a.rec2, a.rec3)
	}
}
