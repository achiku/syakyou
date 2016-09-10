package algorithm

import "testing"

func TestQuickFind_NewQuickFind(t *testing.T) {
	cases := []int{1, 10, 15}
	for _, c := range cases {
		qf := NewQuickFind(c)
		t.Logf("%+v", qf)
	}
}

func TestQuickFind_Union(t *testing.T) {
	qf := NewQuickFind(10)
	qf.Union(0, 1)
	qf.Union(0, 9)
	qf.Union(1, 9)
	qf.Union(0, 8)
	t.Logf("%+v", qf)
}
