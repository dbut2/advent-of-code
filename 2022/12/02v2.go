package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"
	"sync"

	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/timer"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	defer timer.Start()()
	t := test.Register(tests, solve)
	t.Expect(1, 29)
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

	c := make(chan *Cell)
	wg := &sync.WaitGroup{}
	go process(wg, c)

	var starts []*Cell
	var end *Cell
	for i, line := range s {
		for j, cell := range line {
			grid[i][j].MinToEnd = -1
			grid[i][j].Processing = false
			grid[i][j].val = int(cell)

			if string(cell) == "S" {
				grid[i][j].val = 97
			}
			if string(cell) == "E" {
				grid[i][j].MinToEnd = 0
				grid[i][j].val = 122
				end = grid[i][j]
			}

			if grid[i][j].val == 97 {
				starts = append(starts, grid[i][j])
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

	send(wg, c, end)
	wg.Wait()

	min := -1
	for _, start := range starts {
		if min == -1 || start.MinToEnd < min && start.MinToEnd != -1 {
			min = start.MinToEnd
		}

	}
	return min
}

type Grid [][]*Cell

type Cell struct {
	val        int
	Neighbours []*Cell
	MinToEnd   int
	Processing bool
}

func send(wg *sync.WaitGroup, c chan *Cell, cell *Cell) {
	if cell.Processing {
		return
	}
	wg.Add(1)
	cell.Processing = true
	go func() { c <- cell }()
}

func process(wg *sync.WaitGroup, c chan *Cell) {
	for {
		cell := <-c
		neighbours := processCell(cell)
		for _, n := range neighbours {
			send(wg, c, n)
		}
		wg.Done()
	}
}

func processCell(cell *Cell) []*Cell {
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
		return cell.Neighbours
	}
	return nil
}
