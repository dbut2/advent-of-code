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
	h.Expect(1, 51)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	grid := space.NewGridFromInput(s)

	var startPositions [][2][2]int
	for i := 0; i < len(grid); i++ {
		startPositions = append(startPositions, [2][2]int{{i, 0}, {0, 1}})
		startPositions = append(startPositions, [2][2]int{{i, len(grid[0]) - 1}, {0, -1}})
	}
	for i := 0; i < len(grid[0]); i++ {
		startPositions = append(startPositions, [2][2]int{{0, i}, {1, 0}})
		startPositions = append(startPositions, [2][2]int{{len(grid) - 1, i}, {-1, 0}})
	}

	maxSeen := 0
	for _, startPosition := range startPositions {
		beams := [][2][2]int{startPosition}

		seenCoords := sets.Set[[2]int]{}
		seenCoords.Add(startPosition[0])

		seenBeams := sets.SetFrom(beams)

		for {
			var newBeams [][2][2]int
			for _, beam := range beams {
				coords := beam[0]
				direction := beam[1]

				var newDirections [][2]int
				switch grid[coords[0]][coords[1]] {
				case '/':
					newDirections = [][2]int{{-direction[1], -direction[0]}}
				case '\\':
					newDirections = [][2]int{{direction[1], direction[0]}}
				case '-':
					if math.Abs(direction[0]) == 1 {
						newDirections = [][2]int{direction}
					} else {
						newDirections = [][2]int{{1, 0}, {-1, 0}}
					}
				case '|':
					if math.Abs(direction[1]) == 1 {
						newDirections = [][2]int{direction}
					} else {
						newDirections = [][2]int{{0, 1}, {0, -1}}
					}
				case '.':
					newDirections = [][2]int{direction}
				}

				for _, newDirection := range newDirections {
					newCoords := [2]int{coords[0] + newDirection[0], coords[1] + newDirection[1]}

					if !grid.Inside(newCoords[0], newCoords[1]) {
						continue
					}

					if !seenCoords.Has(newCoords) {
						seenCoords.Add(newCoords)
					}

					newBeam := [2][2]int{newCoords, newDirection}
					if seenBeams.Has(newBeam) {
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

		maxSeen = max(maxSeen, len(seenCoords))
	}

	return maxSeen
}
