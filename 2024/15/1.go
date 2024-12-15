package main

import (
	"slices"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input [2]string) int {
	g := space.NewGridFromInput(strings.Split(input[0], "\n"))

	robot, _ := g.Find(func(cell space.Cell, u uint8) bool {
		return u == '@'
	})

	moveMap := map[byte]space.Direction{'^': space.Up, '>': space.Right, 'v': space.Down, '<': space.Left}

	for _, move := range input[1] {
		if move == '\n' {
			continue
		}
		dir := moveMap[byte(move)]

		// find all cells including robot that will need to move
		var toMove []space.Cell
		for cell := robot; ; cell = cell.Move(dir) {
			v := g.Get(cell)
			if *v == '#' {
				toMove = nil
				break
			}
			if *v == '.' {
				break
			}
			toMove = append(toMove, cell)
		}

		// hit wall, cancel
		if len(toMove) == 0 {
			continue
		}

		// parse in filo order
		slices.Reverse(toMove)
		for _, cell := range toMove {
			g.Set(cell.Move(dir), *g.Get(cell))
			g.Set(cell, '.')
		}
		robot = robot.Move(dir)
	}

	total := 0
	for cell, val := range g.Cells() {
		if *val == 'O' {
			total += cell[0] + 100*cell[1]
		}
	}
	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 2028)
	h.Expect(2, 10092)
	h.Run()
}
