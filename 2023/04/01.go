package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/math"
	strings2 "github.com/dbut2/advent-of-code/pkg/strings"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 13)
	fmt.Println(solve(input))
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
