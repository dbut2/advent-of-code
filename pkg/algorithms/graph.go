package algorithms

import "github.com/dbut2/advent-of-code/pkg/sets"

type Graph[T comparable] struct {
	Nodes []T
	Edges []Edge[T]
}

type Edge[T any] struct {
	A, B     T
	Distance int
}

func (g Graph[T]) optimal() (sets.Set[T], map[T]map[T]int) {
	nodeSet := sets.SetFrom(g.Nodes)
	edgeMap := map[T]map[T]int{}
	for _, edge := range g.Edges {
		if edgeMap[edge.A] == nil {
			edgeMap[edge.A] = map[T]int{}
		}
		edgeMap[edge.A][edge.B] = edge.Distance
	}
	return nodeSet, edgeMap
}
