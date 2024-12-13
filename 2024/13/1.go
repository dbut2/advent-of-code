package main

import (
	"math"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/strings"
)

func solve(input [][]string) int {
	tokens := 0
	for _, machine := range input {
		var a, b, p [2]int

		for i, line := range machine {
			nums := strings.Ints(line)
			switch i {
			case 0:
				a[0], a[1] = nums[0], nums[1]
			case 1:
				b[0], b[1] = nums[0], nums[1]
			case 2:
				p[0], p[1] = nums[0], nums[1]
			}
		}

		an := float64(b[0]*p[1]-p[0]*b[1]) / float64(b[0]*a[1]-a[0]*b[1])
		bn := float64(a[0]*p[1]-p[0]*a[1]) / float64(a[0]*b[1]-b[0]*a[1])

		if math.Mod(an, 1) == 0 && math.Mod(bn, 1) == 0 {
			tokens += int(an)*3 + int(bn)
		}
	}

	return tokens
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 480)
	h.Run()
}
