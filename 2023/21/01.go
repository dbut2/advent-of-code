package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
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
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	grid := space.NewGridFromInput(s)
	start := [2]int{}

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 'S' {
				start = [2]int{x, y}
				grid[x][y] = '.'
			}
		}
	}

	seen := map[[2]int][][2]int{}

	possibilites := sets.SetOf(start)
	for i := 0; i < 64; i++ {
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

	return len(possibilites)
}
