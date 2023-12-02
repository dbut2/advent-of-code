package main

import (
	"embed"
	_ "embed"
	"fmt"
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
	t.Expect(1, 8)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	total := 0
	for i, line := range s {
		game := strings.Split(line, ": ")
		rolls := strings.Split(game[1], "; ")

		var maxCubeCounts = make(map[string]int)

		for _, roll := range rolls {
			cubes := strings.Split(roll, ", ")
			for _, cube := range cubes {
				cc := strings.Split(cube, " ")
				count := sti.Sti(cc[0])
				color := cc[1]

				lastMax, _ := maxCubeCounts[color]
				maxCubeCounts[color] = max(lastMax, count)
			}
		}

		//12 red cubes, 13 green cubes, and 14 blue cubes.

		count, _ := maxCubeCounts["red"]
		if count > 12 {
			continue
		}

		count, _ = maxCubeCounts["green"]
		if count > 13 {
			continue
		}

		count, _ = maxCubeCounts["blue"]
		if count > 14 {
			continue
		}

		total += i + 1
	}

	return total
}
