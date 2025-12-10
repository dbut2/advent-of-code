package main

import (
	"math"
	"strings"

	"github.com/draffensperger/golp"

	"github.com/dbut2/advent-of-code/pkg/harness"
	. "github.com/dbut2/advent-of-code/pkg/std"
)

func solve(input []string) int {
	total := 0
	for _, line := range input {
		parts := strings.Split(line, " ")
		targetStr := parts[len(parts)-1]
		buttonsStrs := parts[1 : len(parts)-1]

		targets := Ints(targetStr)
		var affects [][]int
		for _, buttonStr := range buttonsStrs {
			affects = append(affects, Ints(buttonStr))
		}

		lp := golp.NewLP(0, len(buttonsStrs))
		obj := make([]float64, len(buttonsStrs))
		for j := 0; j < len(buttonsStrs); j++ {
			obj[j] = 1.0
			lp.SetInt(j, true)
		}
		lp.SetObjFn(obj)

		for i := range len(targets) {
			row := make([]float64, len(buttonsStrs))
			hasConstraint := false
			for j := range len(buttonsStrs) {
				for _, ci := range affects[j] {
					if ci == i {
						row[j] = 1
						hasConstraint = true
						break
					}
				}
			}
			if hasConstraint {
				lp.AddConstraint(row, golp.EQ, float64(targets[i]))
			}
		}

		lp.Solve()

		vars := lp.Variables()
		for j := 0; j < len(buttonsStrs); j++ {
			total += int(math.Round(vars[j]))
		}
	}
	return total
}

// CGO_CFLAGS="-I/opt/homebrew/Cellar/lp_solve/5.5.2.14/include" CGO_LDFLAGS="-L/opt/homebrew/Cellar/lp_solve/5.5.2.14/lib -llpsolve55" go run 2.go
func main() {
	h := harness.New(solve)
	h.Expect(1, 33)
	h.Run()
}
