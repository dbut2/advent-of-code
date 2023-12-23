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
	visited := sets.Set[T]{}
	var dfs func(T, int) int
	dfs = func(node T, length int) int {
		if node == b {
			return length
		}
		minLength := math.MaxInt
		visited.Add(node)
		for nextNode, nextLength := range g.Edges[node] {
			if !visited.Contains(nextNode) {
				minLength = min(minLength, dfs(nextNode, length+nextLength))
			}
		}
		visited.Remove(node)
		return minLength
	}
	return dfs(a, 0)
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
