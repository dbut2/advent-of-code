package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expected(2, 36)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)
	visits := sets.Set[[2]int]{}

	knots := lists.Fill2D(10, 2, 0)

	visits.Add([2]int{knots[len(knots)-1][0], knots[len(knots)-1][1]})

	for _, line := range s {
		m := strings.Split(line, " ")
		dir := m[0]
		amt := sti.Sti(m[1])

		for i := 0; i < amt; i++ {
			switch dir {
			case "U":
				knots[0][0]--
			case "D":
				knots[0][0]++
			case "L":
				knots[0][1]--
			case "R":
				knots[0][1]++
			}

			for j := 1; j < len(knots); j++ {
				if math.Abs(knots[j-1][0]-knots[j][0]) <= 1 && math.Abs(knots[j-1][1]-knots[j][1]) <= 1 {
					continue
				}

				knots[j][0] += math.Sign(knots[j-1][0] - knots[j][0])
				knots[j][1] += math.Sign(knots[j-1][1] - knots[j][1])
			}

			visits.Add([2]int{knots[len(knots)-1][0], knots[len(knots)-1][1]})
		}
	}

	return len(visits)
}
