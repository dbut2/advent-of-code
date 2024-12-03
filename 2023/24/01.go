package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Run()
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

		x, y, z := sti.Int(splits[0]), sti.Int(splits[1]), sti.Int(splits[2])
		dx, dy, dz := sti.Int(splits[4]), sti.Int(splits[5]), sti.Int(splits[6])

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

			// Create an equation that represents all possible placings of x and y
			// Use the interaction of the two equations to find if inside collision area if at all

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

	// All pairings found twice, return half
	return total / 2
}
