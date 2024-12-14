package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
)

func solve(input [][]int) int {
	type robot struct {
		x, y   int
		dx, dy int
	}

	var rs []robot
	for _, line := range input {
		rs = append(rs, robot{
			x:  line[0],
			y:  line[1],
			dx: line[2],
			dy: line[3],
		})
	}

	width := 0
	height := 0
	for _, r := range rs {
		width = max(width, r.x+1)
		height = max(height, r.y+1)
	}

	for range 100 {
		for i := range rs {
			rs[i].x = (rs[i].x + rs[i].dx + width) % width
			rs[i].y = (rs[i].y + rs[i].dy + height) % height
		}
	}

	quads := [4]int{}
	for _, r := range rs {
		left := r.x < width/2
		right := r.x > width/2

		up := r.y < height/2
		down := r.y > height/2

		switch {
		case left && up:
			quads[0]++
		case left && down:
			quads[1]++
		case right && up:
			quads[2]++
		case right && down:
			quads[3]++
		}
	}

	return quads[0] * quads[1] * quads[2] * quads[3]
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 12)
	h.Run()
}
