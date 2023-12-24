package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Solve()
}

type hailstone struct {
	px, py int
	vx, vy int
}

func solve(input string) int {
	s := utils.ParseInput(input)

	hs := sets.Set[hailstone]{}

	for j, line := range s {
		_, _ = j, line

		line = strings.ReplaceAll(line, ",", "")

		splits := strings.Split(line, " ")

		x, y, z := sti.Sti(splits[0]), sti.Sti(splits[1]), sti.Sti(splits[2])
		dx, dy, dz := sti.Sti(splits[4]), sti.Sti(splits[5]), sti.Sti(splits[6])

		hs.Add(hailstone{
			px: x,
			py: y,
			vx: dx,
			vy: dy,
		})

		_, _ = z, dz
	}

	collisionAreaX := [2]float64{200000000000000, 400000000000000}
	collisionAreaY := [2]float64{200000000000000, 400000000000000}

	total := 0

	for h1 := range hs {
		for h2 := range hs {
			if h1 == h2 {
				continue
			}

			y1 := float64(h1.py)
			m1 := float64(h1.vy) / float64(h1.vx)
			x1 := float64(h1.px)
			c1 := y1 - m1*x1

			y2 := float64(h2.py)
			m2 := float64(h2.vy) / float64(h2.vx)
			x2 := float64(h2.px)
			c2 := y2 - m2*x2

			if m1 == m2 {
				continue
			}

			x := (c2 - c1) / (m1 - m2)
			y := m1*x + c1

			if x < collisionAreaX[0] || x > collisionAreaX[1] {
				continue
			}
			if y < collisionAreaY[0] || y > collisionAreaY[1] {
				continue
			}

			if (x-x1)/float64(h1.vx) < 0 || (x-x2)/float64(h2.vx) < 0 {
				continue
			}

			total++
		}
	}

	return total / 2
}
