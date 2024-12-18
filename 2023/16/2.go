package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 51)
	h.Run()
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

		maxSeen = max(maxSeen, len(seenCoords))
	}

	return maxSeen
}
