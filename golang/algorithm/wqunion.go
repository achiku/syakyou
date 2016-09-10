package algorithm

import (
	"errors"
	"fmt"
)

// WeightedQuickUnion weighted quick union
type WeightedQuickUnion struct {
	ids  []int
	size []int
}

// NewWeightedQuickUnion create new weight quick union
func NewWeightedQuickUnion(n int) WeightedQuickUnion {
	var ids []int
	var size []int
	for i := 0; i <= n; i++ {
		ids = append(ids, i)
		size = append(size, 1)
	}
	return WeightedQuickUnion{
		ids:  ids,
		size: size,
	}
}

func (qu *WeightedQuickUnion) findRoot(x int) int {
	parent := qu.ids[x]
	if parent == x {
		return parent
	}
	return qu.findRoot(parent)
}

// Union connect p and q
func (qu *WeightedQuickUnion) Union(p int, q int) error {
	if p < 0 || q < 0 {
		return errors.New("p and q must be greater or equal to 0")
	}
	if p > len(qu.ids) || q > len(qu.ids) {
		return fmt.Errorf("p and q must be smaller or equal to %d", len(qu.ids))
	}
	pRoot := qu.findRoot(p)
	qRoot := qu.findRoot(q)
	if pRoot == qRoot {
		return nil
	}
	if qu.size[pRoot] < qu.size[qRoot] {
		qu.ids[pRoot] = qRoot
		qu.size[qRoot] = +qu.size[pRoot]
	} else {
		qu.ids[qRoot] = pRoot
		qu.size[pRoot] = +qu.size[qRoot]
	}
	return nil
}

// Connected check if p and q are connected
func (qu *WeightedQuickUnion) Connected(p int, q int) bool {
	if qu.findRoot(p) == qu.findRoot(q) {
		return true
	}
	return false
}
