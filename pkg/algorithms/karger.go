package algorithms

import (
	"math/rand"

	"github.com/dbut2/advent-of-code/pkg/sets"
)

func Karger[T comparable](g Graph[T]) ([]Edge[T], [2][]T) {
	nodeSet := sets.SetFrom(g.Nodes)
	edgeSet := sets.Set[*Edge[T]]{}
	oldEdges := map[*Edge[T]]Edge[T]{}
	for _, edge := range g.Edges {
		newEdge := edge
		edgeSet.Add(&newEdge)
		oldEdges[&newEdge] = edge
	}
	collapsedInto := map[T][]T{}

	for len(nodeSet) > 2 {
		i := rand.Intn(len(edgeSet))
		edge := edgeSet.Slice()[i]
		collapsedInto[edge.A] = append(collapsedInto[edge.A], edge.B)
		collapsedInto[edge.A] = append(collapsedInto[edge.A], collapsedInto[edge.B]...)
		edgeSet.Remove(edge)
		nodeSet.Remove(edge.B)
		for edge2 := range edgeSet {
			if edge2.A == edge.B {
				edge2.A = edge.A
			}
			if edge2.B == edge.B {
				edge2.B = edge.A
			}
			if edge2.A == edge2.B {
				edgeSet.Remove(edge2)
			}
		}
	}

	retEdges := make([]Edge[T], 0, len(edgeSet))
	for edge := range edgeSet {
		retEdges = append(retEdges, oldEdges[edge])
	}

	nodeSlice := nodeSet.Slice()
	retNodePools := [2][]T{}
	retNodePools[0] = append(collapsedInto[nodeSlice[0]], nodeSlice[0])
	retNodePools[1] = append(collapsedInto[nodeSlice[1]], nodeSlice[1])

	return retEdges, retNodePools
}
