package main

import "fmt"

// NodeRec struct
type NodeRec struct {
	ID               int
	Name             string
	HasAssociatedRec bool
	AssociatedID     int
}

// Node struct
type Node struct {
	ID            int
	Name          string
	AssociatedIDs []int
}

// CreateNode create node
func CreateNode(recs []NodeRec) []Node {
	var node []Node
	for i, r := range recs {
		fmt.Printf("%d: %+v\n", i, r)
		if r.HasAssociatedRec {
		}
	}
	return node
}

func main() {
	recs := []NodeRec{
		NodeRec{ID: 1, Name: "a", HasAssociatedRec: false},
		NodeRec{ID: 2, Name: "b", HasAssociatedRec: false},
		NodeRec{ID: 3, Name: "ba", HasAssociatedRec: true, AssociatedID: 2},
		NodeRec{ID: 4, Name: "aa", HasAssociatedRec: true, AssociatedID: 1},
		NodeRec{ID: 5, Name: "bb", HasAssociatedRec: true, AssociatedID: 2},
		NodeRec{ID: 6, Name: "bc", HasAssociatedRec: true, AssociatedID: 2},
	}

	for _, r := range recs {
		fmt.Printf("NodeRec: %+v\n", r)
	}

	for _, r := range CreateNode(recs) {
		fmt.Printf("Node: %+v\n", r)
	}
}
