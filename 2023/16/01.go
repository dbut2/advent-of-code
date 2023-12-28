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
	h.Expect(1, 46)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	grid := space.NewGridFromInput(s)

	beams := [][2][2]int{{{0, 0}, {1, 0}}}

	seenCoords := sets.Set[[2]int]{}
	seenCoords.Add([2]int{0, 0})
	seenBeams := sets.SetFrom(beams)
	for {
		var newBeams [][2][2]int
		for _, beam := range beams {
			coords := beam[0]
			direction := beam[1]

			var newDirections []space.Direction
			switch grid[coords[0]][coords[1]] {
			case '/':
				newDirections = []space.Direction{{-direction[1], -direction[0]}}
			case '\\':
				newDirections = []space.Direction{{direction[1], direction[0]}}
			case '-':
				if math.Abs(direction[0]) == 1 {
					newDirections = []space.Direction{direction}
				} else {
					newDirections = []space.Direction{space.Left, space.Right}
				}
			case '|':
				if math.Abs(direction[1]) == 1 {
					newDirections = []space.Direction{direction}
				} else {
					newDirections = []space.Direction{space.Up, space.Down}
				}
			case '.':
				newDirections = []space.Direction{direction}
			}

			for _, newDirection := range newDirections {
				newCoords := space.Cell.Move(coords, newDirection)

				if !grid.Inside(newCoords) {
					continue
				}

				if !seenCoords.Contains(newCoords) {
					seenCoords.Add(newCoords)
				}

				newBeam := [2][2]int{newCoords, newDirection}
				if seenBeams.Contains(newBeam) {
					continue
				}
				seenBeams.Add(newBeam)
				newBeams = append(newBeams, newBeam)
			}
		}

		if len(newBeams) == 0 {
			break
		}

		beams = newBeams
	}

	return len(seenCoords)
}
