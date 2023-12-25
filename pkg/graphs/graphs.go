package graphs

import (
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"math/rand"
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

type Edge[T any] struct {
	A, B     T
	Distance int
}

func (g *Graph[T]) Karger() ([]Edge[T], [2][]T) {
	nodes := sets.Set[T]{}
	edges := sets.Set[*Edge[T]]{}
	oldEdges := map[*Edge[T]]Edge[T]{}
	for in, outs := range g.Edges {
		nodes.Add(in)
		for out, distance := range outs {
			nodes.Add(out)
			e := Edge[T]{
				A:        in,
				B:        out,
				Distance: distance,
			}
			edges.Add(&e)
			oldEdges[&e] = e
		}
	}

	collapsedInto := map[T][]T{}

	for len(nodes) > 2 {
		i := rand.Intn(len(edges))
		edge := edges.Slice()[i]

		collapsedInto[edge.A] = append(collapsedInto[edge.A], edge.B)
		collapsedInto[edge.A] = append(collapsedInto[edge.A], collapsedInto[edge.B]...)

		edges.Remove(edge)
		nodes.Remove(edge.B)
		for edge2 := range edges {
			if edge2.A == edge.B {
				edge2.A = edge.A
			}
			if edge2.B == edge.B {
				edge2.B = edge.A
			}
			if edge2.A == edge2.B {
				edges.Remove(edge2)
			}
		}
	}

	retEdges := make([]Edge[T], 0, len(edges))
	for edge := range edges {
		retEdges = append(retEdges, oldEdges[edge])
	}

	nodeSlice := nodes.Slice()
	retNodePools := [2][]T{}
	retNodePools[0] = append(collapsedInto[nodeSlice[0]], nodeSlice[0])
	retNodePools[1] = append(collapsedInto[nodeSlice[1]], nodeSlice[1])

	return retEdges, retNodePools
}
