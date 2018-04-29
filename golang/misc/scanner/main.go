package main

import "strings"

var data = `100,data1,100data
101,data2,101data
102,data3,102data
100,data4,100data
101,data5,101data
100,data6,100data
100,data7,100data
200,data8,200data
100,data9,100data
101,data10,101data
`

type parser struct {
	f file
}

func newParser() *parser {
	var f file
	return &parser{
		f: f,
	}
}

func (p *parser) parse(l string) {
	d := strings.Split(l, ",")
	rec := record{
		category:     d[0],
		content:      d[1],
		categoryData: d[2],
	}
	// if category ends with 0, it's a start of the new transaction
	// others are component of the transaction
	// category must be ascending order within the same transaction
	typ := rec.categoryType()
	switch typ {
	case "0":
		tran := &transaction{
			rec1: &rec,
		}
		p.f = append(p.f, tran)
	case "1":
		t := p.f[len(p.f)-1]
		t.rec2 = &rec
	case "2":
		t := p.f[len(p.f)-1]
		t.rec3 = &rec
	}
}

type transaction struct {
	rec1 *record `json:"omitempty,rec_1"`
	rec2 *record `json:"omitempty,rec_2"`
	rec3 *record `json:"omitempty,rec_3"`
}

type record struct {
	category     string `json:"category"`
	content      string `json:"content"`
	categoryData string `json:"category_data"`
}

func (r *record) categoryType() string {
	return string(r.category[2])
}

type file []*transaction
