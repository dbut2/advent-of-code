package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

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
	t.Expect(1, 13)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)
	visits := sets.Set[[2]int]{}

	hx, hy := 0, 0
	tx, ty := 0, 0

	visits.Add([2]int{tx, ty})

	for _, line := range s {
		m := strings.Split(line, " ")
		dir := m[0]
		amt := sti.Sti(m[1])

		for i := 0; i < amt; i++ {
			switch dir {
			case "U":
				hx--
			case "D":
				hx++
			case "L":
				hy--
			case "R":
				hy++
			}

			if math.Abs(hx-tx) <= 1 && math.Abs(hy-ty) <= 1 {
				continue
			}

			tx += math.Sign(hx - tx)
			ty += math.Sign(hy - ty)

			visits.Add([2]int{tx, ty})
		}
	}

	return len(visits)
}
