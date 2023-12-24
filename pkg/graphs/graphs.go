package graphs

import (
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sets"
)

type Graph[T comparable] struct {
	Nodes sets.Set[T]
	Edges map[T]map[T]int
}

func New[T comparable]() Graph[T] {
	g := Graph[T]{}
	g.ensure()
	return g
}

func (g *Graph[T]) ensure() {
	if g.Nodes == nil {
		g.Nodes = sets.Set[T]{}
	}
	if g.Edges == nil {
		g.Edges = map[T]map[T]int{}
	}
}

func (g *Graph[T]) Connect(a, b T, distance int) {
	g.ensure()
	if !g.Nodes.Contains(a) {
		g.Nodes.Add(a)
		g.Edges[a] = map[T]int{}
	}
	if !g.Nodes.Contains(b) {
		g.Nodes.Add(b)
		g.Edges[b] = map[T]int{}
	}
	g.Edges[a][b] = distance
}

func (g *Graph[T]) Minimise(a, b T) int {
	distances := make(map[T]int)
	for node := range g.Nodes {
		distances[node] = math.MaxInt
	}
	distances[a] = 0

	visited := sets.Set[T]{}

	for len(visited) < len(g.Nodes) {
		current := a
		minDistance := math.MaxInt
		for node, distance := range distances {
			if !visited.Contains(node) && distance < minDistance {
				current = node
				minDistance = distance
			}
		}

		visited.Add(current)

		if current == b {
			return distances[current]
		}

		for node, distance := range g.Edges[current] {
			if !visited.Contains(node) {
				newDistance := distances[current] + distance
				distances[node] = min(distances[node], newDistance)
			}
		}
	}

	panic("no path")
}

func (g *Graph[T]) Maximise(a, b T) int {
	visited := sets.Set[T]{}
	var dfs func(T, int) int
	dfs = func(node T, length int) int {
		if node == b {
			return length
		}
		maxLength := math.MinInt
		visited.Add(node)
		for nextNode, nextLength := range g.Edges[node] {
			if !visited.Contains(nextNode) {
				maxLength = max(maxLength, dfs(nextNode, length+nextLength))
			}
		}
		visited.Remove(node)
		return maxLength
	}
	return dfs(a, 0)
}
