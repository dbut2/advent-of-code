package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	grid := newGridFromInput(s)
	start := [2]int{}

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 'S' {
				start = [2]int{x, y}
				grid[x][y] = '.'
			}
		}
	}

	n := 26501365
	seen := map[[2]int][][2]int{}

	// visit only the first n%262 + 262 steps
	// we can use this to extrapolate a larger visit
	possibilites := sets.SetOf(start)
	for i := 0; i < n%262+262; i++ {
		nextPosibilites := sets.Set[[2]int]{}
		for coord := range possibilites {
			if len(seen[coord]) > 0 {
				for _, nextCoord := range seen[coord] {
					nextPosibilites.Add(nextCoord)
				}
				continue
			}
			next := [][2]int{}
			for nextCoord, cell := range grid.Adjacent(coord) {
				if *cell == '.' {
					next = append(next, nextCoord)
					nextPosibilites.Add(nextCoord)
				}
			}
			seen[coord] = append(seen[coord], next...)
		}
		possibilites = nextPosibilites
	}

	// group each 1x1 cell into a 131x131 cell
	gridCounts := map[[2]int][][2]int{}
	for possibility := range possibilites {
		x := (possibility[0]) / 131
		if possibility[0] < 0 {
			n := possibility[0]
			i := 0
			for n < 0 {
				n += 131
				i--
			}
			x = i
		}
		y := (possibility[1]) / 131
		if possibility[1] < 0 {
			n := possibility[1]
			i := 0
			for n < 0 {
				n += 131
				i--
			}
			y = i
		}
		gridCounts[[2]int{x, y}] = append(gridCounts[[2]int{x, y}], possibility)
	}

	r := n / 131
	o := math.Pow(r-((r+1)%2), 2) // visits on odd grids
	e := math.Pow(r-(r%2), 2)     // visits on even grids

	// Use the grid to extrapolate the number of visits for each variation of possible grid visit states
	total := 0
	total += len(gridCounts[[2]int{0, 0}]) * o                   // visits on fully visited odd grids
	total += len(gridCounts[[2]int{1, 0}]) * e                   // visits on fully visited even grids
	total += len(gridCounts[[2]int{2, 0}])                       // visits on easternmost grid
	total += len(gridCounts[[2]int{-2, 0}])                      // visits on westernmost grid
	total += len(gridCounts[[2]int{0, 2}])                       // visits on northernmost grid
	total += len(gridCounts[[2]int{0, -2}])                      // visits on southernmost grid
	total += len(gridCounts[[2]int{1, 1}]) * (((n) / 131) - 1)   // visits on inner of partially visited grids along north-eastern edge
	total += len(gridCounts[[2]int{-1, 1}]) * (((n) / 131) - 1)  // north-western
	total += len(gridCounts[[2]int{1, -1}]) * (((n) / 131) - 1)  // south-eastern
	total += len(gridCounts[[2]int{-1, -1}]) * (((n) / 131) - 1) // south-western
	total += len(gridCounts[[2]int{2, 1}]) * ((n) / 131)         // visits on outer of partially visited grids along north-eastern edge
	total += len(gridCounts[[2]int{-2, 1}]) * ((n) / 131)        // north-western
	total += len(gridCounts[[2]int{2, -1}]) * ((n) / 131)        // south-eastern
	total += len(gridCounts[[2]int{-2, -1}]) * ((n) / 131)       // south-western

	return total
}

type infiniteGrid[T any] space.Grid[T]

func (i *infiniteGrid[T]) Grid() *space.Grid[T] {
	return (*space.Grid[T])(i)
}

func newGridFromInput(s []string) infiniteGrid[uint8] {
	return infiniteGrid[uint8](space.NewGridFromInput(s))
}

func (i *infiniteGrid[T]) Get(cell space.Cell) *T {
	return i.Grid().Get(i.Wrap(cell))
}

func (i *infiniteGrid[T]) Set(cell space.Cell, v T) {
	i.Grid().Set(i.Wrap(cell), v)
}

func (i *infiniteGrid[T]) Adjacent(cell space.Cell) map[space.Cell]*T {
	return map[space.Cell]*T{
		cell.Move(space.North): i.Get(cell.Move(space.North)),
		cell.Move(space.South): i.Get(cell.Move(space.South)),
		cell.Move(space.East):  i.Get(cell.Move(space.East)),
		cell.Move(space.West):  i.Get(cell.Move(space.West)),
	}
}

func (i *infiniteGrid[T]) Wrap(cell space.Cell) space.Cell {
	for cell[0] < 0 {
		cell[0] += len(*i)
	}
	cell[0] %= len(*i)
	for cell[1] < 0 {
		cell[1] += len(*i)
	}
	cell[1] %= len(*i)
	return cell
}
