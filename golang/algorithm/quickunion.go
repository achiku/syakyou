package algorithm

import (
	"errors"
	"fmt"
)

// QuickUnion quick union
type QuickUnion struct {
	ids []int
}

// NewQuickUnion create new quick union
func NewQuickUnion(n int) QuickUnion {
	var ids []int
	for i := 0; i <= n; i++ {
		ids = append(ids, i)
	}
	return QuickUnion{ids: ids}
}

func (qu *QuickUnion) findRoot(x int) int {
	parent := qu.ids[x]
	if parent == x {
		return parent
	}
	return qu.findRoot(parent)
}

// Union connect p and q
func (qu *QuickUnion) Union(p int, q int) error {
	if p < 0 || q < 0 {
		return errors.New("p and q must be greater or equal to 0")
	}
	if p > len(qu.ids) || q > len(qu.ids) {
		return fmt.Errorf("p and q must be smaller or equal to %d", len(qu.ids))
	}
	pRoot := qu.findRoot(p)
	qRoot := qu.findRoot(q)
	qu.ids[pRoot] = qRoot
	return nil
}

// Connected check if p and q are connected
func (qu *QuickUnion) Connected(p int, q int) bool {
	if qu.findRoot(p) == qu.findRoot(q) {
		return true
	}
	return false
}
