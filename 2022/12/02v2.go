package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"
	"sync"

	"github.com/dbut2/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 29)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")

	grid := Grid{}
	for i := range s {
		line := Line{}
		for range s[i] {
			line = append(line, &Cell{})
		}
		grid = append(grid, line)
	}

	rows := len(grid)
	cols := len(grid[0])

	// fill each cell with data and identify possible starts and end cells
	var starts []*Cell
	var end *Cell
	for i, line := range s {
		for j, rawCell := range line {
			cell := grid[i][j]

			if rawCell == 83 {
				rawCell = 97
				s[i] = s[i][:j] + "a" + s[i][j+1:]
			}
			if rawCell == 69 {
				cell.MinToEnd = 0
				cell.Seen = true
				end = cell
				s[i] = s[i][:j] + "z" + s[i][j+1:]
				rawCell = 122
			}

			if rawCell == 97 {
				starts = append(starts, grid[i][j])
			}
		}
	}

	// fill neighbors field for neighbors that can move into current cell
	for i, row := range grid {
		for j, cell := range row {
			cval := s[i][j]

			if i > 0 {
				neighbor := grid[i-1][j]
				nval := s[i-1][j]

				diff := cval - nval
				if diff <= 1 || diff >= 128 {
					cell.Neighbors[0] = neighbor
				}
			}
			if i < rows-1 {
				neighbor := grid[i+1][j]
				nval := s[i+1][j]

				diff := cval - nval
				if diff <= 1 || diff >= 128 {
					cell.Neighbors[1] = neighbor
				}
			}
			if j > 0 {
				neighbor := grid[i][j-1]
				nval := s[i][j-1]

				diff := cval - nval
				if diff <= 1 || diff >= 128 {
					cell.Neighbors[2] = neighbor
				}
			}
			if j < cols-1 {
				neighbor := grid[i][j+1]
				nval := s[i][j+1]

				diff := cval - nval
				if diff <= 1 || diff >= 128 {
					cell.Neighbors[3] = neighbor
				}
			}
		}
	}

	wg := &sync.WaitGroup{}
	send(wg, end)
	wg.Wait()

	min := 0
	for _, start := range starts {
		if start.MinToEnd == 0 {
			continue
		}
		if min == 0 || start.MinToEnd < min {
			min = start.MinToEnd
		}
	}

	return min
}

type Grid []Line
type Line []*Cell

type Cell struct {
	Neighbors  [4]*Cell
	MinToEnd   int
	Seen       bool
	Processing bool
}

// send cell to chan, manages wg and cell.Processing for efficiency
func send(wg *sync.WaitGroup, cell *Cell) {
	if cell.Processing {
		return
	}
	cell.Processing = true
	wg.Add(1)
	go process(wg, cell)
}

// process will listen on channel and process any cells it sees
func process(wg *sync.WaitGroup, cell *Cell) {
	cell.Processing = false
	for _, n := range cell.Neighbors {
		if n == nil {
			continue
		}
		nm := cell.MinToEnd + 1
		if !n.Seen {
			n.MinToEnd = nm
			n.Seen = true
			send(wg, n)
			continue
		}
		if nm < n.MinToEnd {
			n.MinToEnd = nm
			send(wg, n)
		}
	}
	wg.Done()
}
