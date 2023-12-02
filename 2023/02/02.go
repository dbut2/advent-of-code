package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/dbut2/advent-of-code/pkg/benchmark"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(2, 2286)
	fmt.Println(solve(input))
	benchmark.Run(func() {
		solve(input)
	}, benchmark.Count(1000))
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
				count := sti.Sti(cc[0])
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
