package algorithms

import (
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sets"
)

func Dijkstra[T comparable](g Graph[T], a, b T) int {
	nodeSet, edgeMap := g.optimal()

	distances := make(map[T]int, len(nodeSet))
	for node := range nodeSet {
		distances[node] = math.MaxInt
	}
	distances[a] = 0

	visited := make(sets.Set[T], len(nodeSet))
	for len(visited) < len(nodeSet) {
		current := a
		minDistance := math.MaxInt
		for node, distance := range distances {
			if !visited.Contains(node) && distance < minDistance {
				current = node
				minDistance = distance
			}
		}
		if current == b {
			return distances[current]
		}
		visited.Add(current)
		for node, distance := range edgeMap[current] {
			if !visited.Contains(node) {
				newDistance := distances[current] + distance
				distances[node] = min(distances[node], newDistance)
			}
		}
	}

	panic("no path")
}
