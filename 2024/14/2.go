package main

import (
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sets"
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

	c := 0
	for {
		c++
		tiles := sets.Set[[2]int]{}
		for i := range rs {
			r := rs[i]
			r.x = (r.x + r.dx + width) % width
			r.y = (r.y + r.dy + height) % height
			rs[i] = r
			tiles.Add([2]int{r.x, r.y})
		}

		// Check grid for any 3x3 filled in square, assuming to be a tree
		hasSquare := false
		for x := range width - 3 {
			for y := range height - 3 {
				thisHas := true
				for dx := range 3 {
					for dy := range 3 {
						if !tiles.Contains([2]int{x + dx, y + dy}) {
							thisHas = false
							break
						}
					}
				}
				if thisHas {
					hasSquare = true
				}
			}
		}
		if hasSquare {
			break
		}
	}
	return c
}

func main() {
	h := harness.New(solve)
	h.Run()
}
