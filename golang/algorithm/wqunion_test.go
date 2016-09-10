package algorithm

import "testing"

func TestWeightedQuickUnion_NewWeightedQuickUnion(t *testing.T) {
	cases := []int{1, 10, 15}
	for _, c := range cases {
		qu := NewWeightedQuickUnion(c)
		t.Logf("%+v", qu)
	}
}

func TestWeightedQuickUnion_Union(t *testing.T) {
	qu := NewWeightedQuickUnion(10)
	qu.Union(7, 1)
	qu.Union(7, 2)
	qu.Union(7, 3)
	qu.Union(0, 9)
	qu.Union(0, 8)
	qu.Union(1, 8)
	t.Logf("%+v", qu)
}
