package main

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

type transaction struct {
	rec1 *record
	rec2 *record
	rec3 *record
}

type record struct {
	category     string
	content      string
	categoryData string
}

func (r *record) categoryType() string {
	return string(r.category[2])
}

type file []*transaction
