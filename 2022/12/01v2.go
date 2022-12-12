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
	t.Expect(1, 31)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")
	grid := Grid{}

	for i := range s {
		var line []*Cell
		for range s[i] {
			line = append(line, &Cell{})
		}
		grid = append(grid, line)
	}

	var start *Cell

	c := make(chan *Cell)
	wg := &sync.WaitGroup{}

	for i, line := range s {
		for j, cell := range line {
			grid[i][j].MinToEnd = -1
			grid[i][j].Processing = false
			grid[i][j].val = int(cell)

			if string(cell) == "S" {
				start = grid[i][j]
				grid[i][j].val = 97
			}
			if string(cell) == "E" {
				grid[i][j].MinToEnd = 0
				grid[i][j].val = 122
				send(wg, c, grid[i][j])
			}
		}
	}

	for i, row := range grid {
		for j, cell := range row {
			nCoords := [][]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
			for _, coord := range nCoords {
				if coord[0] >= 0 && coord[0] < len(grid) && coord[1] >= 0 && coord[1] < len(grid[coord[0]]) {
					neighbor := grid[coord[0]][coord[1]]
					if cell.val-neighbor.val <= 1 {
						cell.Neighbours = append(cell.Neighbours, neighbor)
					}
				}
			}
		}
	}

	process(wg, c)
	return start.MinToEnd
}

type Grid [][]*Cell

type Cell struct {
	val        int
	Neighbours []*Cell
	MinToEnd   int
	Processing bool
}

func send(wg *sync.WaitGroup, c chan *Cell, cell *Cell) {
	wg.Add(1)
	go func() {
		if cell.Processing {
			wg.Done()
			return
		}
		cell.Processing = true
		c <- cell
	}()
}

func process(wg *sync.WaitGroup, c chan *Cell) {
	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()
	for loop := true; loop; {
		select {
		case cell := <-c:
			cell.Processing = false
			changed := false
			for _, n := range cell.Neighbours {
				nm := cell.MinToEnd + 1
				if nm < n.MinToEnd || n.MinToEnd < 0 {
					n.MinToEnd = nm
					changed = true
				}
			}
			if changed {
				for _, n := range cell.Neighbours {
					send(wg, c, n)
				}
			}
			wg.Done()
		case <-done:
			loop = false
		}
	}
}
