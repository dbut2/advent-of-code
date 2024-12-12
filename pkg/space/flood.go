package space

import (
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sets"
)

func FloodCell[T comparable](g Grid[T], cell Cell) []Cell {
	val := g.Get(cell)
	seen := sets.Set[Cell]{}
	seen.Add(cell)
	queue := lists.Queue[Cell]{cell}
	for next := range queue.Seq {
		for c, v := range g.Adjacent(next) {
			if *val == *v {
				if !seen.Contains(c) {
					queue.Push(c)
					seen.Add(c)
				}
			}
		}
	}
	return seen.Slice()
}

func FloodAll[T comparable](g Grid[T]) [][]Cell {
	all := sets.SetFrom(g.FindAll(func(_ Cell, _ T) bool { return true }))
	var floods [][]Cell
	for len(all) > 0 {
		flood := FloodCell(g, all.Slice()[0])
		floods = append(floods, flood)
		for _, cell := range flood {
			all.Remove(cell)
		}
	}
	return floods
}
