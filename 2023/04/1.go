package main

import (
	"embed"
	_ "embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/math"
	strings2 "github.com/dbut2/advent-of-code/pkg/strings"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 13)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	total := 0

	for _, line := range s {
		game := strings.Split(line, ": ")
		games := strings.Split(game[1], "|")

		left := strings2.Ints(games[0])
		right := strings2.Ints(games[1])

		matches := len(lists.Intersection(left, right))
		if matches == 0 {
			continue
		}
		total += math.Pow(2, matches-1)
	}

	return total

}
