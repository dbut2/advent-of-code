package main

import (
	"embed"
	_ "embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"

	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(2, 2286)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	total := 0
	for _, line := range s {
		game := strings.Split(line, ": ")
		rolls := strings.Split(game[1], "; ")

		var maxCubesCounts = make(map[string]int)

		for _, roll := range rolls {
			cubes := strings.Split(roll, ", ")
			for _, cube := range cubes {
				cc := strings.Split(cube, " ")
				count := sti.Int(cc[0])
				color := cc[1]

				lastMax, _ := maxCubesCounts[color]
				maxCubesCounts[color] = max(lastMax, count)
			}
		}

		product := 1
		for _, count := range maxCubesCounts {
			product *= count
		}
		total += product
	}

	return total
}
