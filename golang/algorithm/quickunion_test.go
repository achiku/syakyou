package algorithm

import "testing"

func TestQuickUnion_NewQuickUnion(t *testing.T) {
	cases := []int{1, 10, 15}
	for _, c := range cases {
		qu := NewQuickUnion(c)
		t.Logf("%+v", qu)
	}
}

func TestQuickUnion_Union(t *testing.T) {
	qu := NewQuickUnion(10)
	qu.Union(0, 1)
	qu.Union(0, 9)
	qu.Union(1, 9)
	qu.Union(0, 8)
	t.Logf("%+v", qu)
	t.Logf("%t", qu.Connected(0, 8))
	t.Logf("%t", qu.Connected(5, 8))
	qu.Union(1, 5)
	t.Logf("%t", qu.Connected(5, 8))
}
