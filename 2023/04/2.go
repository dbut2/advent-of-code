package main

import (
	"embed"
	_ "embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	strings2 "github.com/dbut2/advent-of-code/pkg/strings"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(2, 30)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	total := 0
	multipliers := map[int]int{}
	for i := range s {
		multipliers[i] = 1
	}

	for i, line := range s {
		game := strings.Split(line, ": ")
		games := strings.Split(game[1], "|")

		left := strings2.Ints(games[0])
		right := strings2.Ints(games[1])

		matches := lists.Intersection(left, right)
		multiplier := multipliers[i]
		total += multiplier
		for j := range matches {
			multipliers[i+j+1] += multiplier
		}
	}

	return total
}
