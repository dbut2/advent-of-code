package main

import (
	"slices"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
)

type grid = space.Grid[byte]

func solve(input [2]string) int {
	p1 := input[0]
	p1 = strings.ReplaceAll(p1, "#", "##")
	p1 = strings.ReplaceAll(p1, "O", "[]")
	p1 = strings.ReplaceAll(p1, ".", "..")
	p1 = strings.ReplaceAll(p1, "@", "@.")
	g := space.NewGridFromInput(strings.Split(p1, "\n"))

	robot, _ := g.Find(func(cell space.Cell, u uint8) bool {
		return u == '@'
	})

	moveMap := map[byte]space.Direction{'^': space.Up, '>': space.Right, 'v': space.Down, '<': space.Left}

	for _, move := range input[1] {
		if move == '\n' {
			continue
		}
		dir := moveMap[byte(move)]

		var toMove []space.Cell

		// queue check every cell that may need to move
		queue := lists.Queue[space.Cell]{robot}
		seen := sets.Set[space.Cell]{}
		for cell := range queue.Seq {
			if seen.Contains(cell) {
				continue
			}
			seen.Add(cell)

			v := g.Get(cell)
			if *v == '#' {
				toMove = nil
				break
			}
			if *v == '.' {
				continue
			}

			// if partial block, now include other part
			if *v == '[' {
				queue.Push(cell.Move(space.Right))
			}
			if *v == ']' {
				queue.Push(cell.Move(space.Left))
			}
			queue.Push(cell.Move(dir))
			toMove = append(toMove, cell)
		}

		if len(toMove) == 0 {
			continue
		}

		// sort in reverse order of direction
		slices.SortFunc(toMove, func(a, b space.Cell) int {
			return dir[0]*(b[0]-a[0]) + dir[1]*(b[1]-a[1])
		})

		for i := range toMove {
			g.Set(toMove[i].Move(dir), *g.Get(toMove[i]))
			g.Set(toMove[i], '.')
		}
		robot = robot.Move(dir)
	}

	total := 0
	for _, c := range g.FindAll(func(cell space.Cell, u uint8) bool {
		return u == '['
	}) {
		total += c[0] + 100*c[1]
	}
	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(2, 9021)
	h.Run()
}
