package algorithm

import (
	"errors"
	"fmt"
)

// QuickFind quick find
type QuickFind struct {
	ids []int
}

// NewQuickFind create new quick find
func NewQuickFind(n int) QuickFind {
	var ids []int
	for i := 0; i <= n; i++ {
		ids = append(ids, i)
	}
	return QuickFind{ids: ids}
}

// Union union two elements
func (qf *QuickFind) Union(p int, q int) error {
	if p < 0 || q < 0 {
		return errors.New("p and q must be greater or equal to 0")
	}
	if p > len(qf.ids) || q > len(qf.ids) {
		return fmt.Errorf("p and q must be smaller or equal to %d", len(qf.ids))
	}
	pid := qf.ids[p]
	qid := qf.ids[q]

	// p and q are already connected
	if pid == qid {
		return nil
	}
	// p and q are already connected
	for i, v := range qf.ids {
		if v == pid {
			qf.ids[i] = qid
		}
	}
	return nil
}
