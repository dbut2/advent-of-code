package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(1, 5)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	a, b, c := 0, 0, 0

	for j, line := range s {
		_, _ = j, line

		parts := strings.Split(line, "~")
		for _, part := range parts {
			split := strings.Split(part, ",")
			a = max(a, sti.Sti(split[0]))
			b = max(b, sti.Sti(split[1]))
			c = max(c, sti.Sti(split[2]))
		}
	}

	type part struct {
		blocks [][3]int
	}

	parts := []*part{}
	cube := space.NewCube[*part](a+1, b+1, c+1)

	for j, line := range s {
		_, _ = j, line

		p := &part{blocks: make([][3]int, 0)}
		parts = append(parts, p)

		coords := [3][2]int{}

		parts := strings.Split(line, "~")
		for i, part := range parts {
			split := strings.Split(part, ",")
			coords[0][i] = sti.Sti(split[0])
			coords[1][i] = sti.Sti(split[1])
			coords[2][i] = sti.Sti(split[2])
		}

		for x := coords[0][0]; x <= coords[0][1]; x++ {
			for y := coords[1][0]; y <= coords[1][1]; y++ {
				for z := coords[2][0]; z <= coords[2][1]; z++ {
					cube[x][y][z] = p
					p.blocks = append(p.blocks, [3]int{x, y, z})
				}
			}
		}
	}

	for {
		fell := false

		for _, p := range parts {
			canFell := true
			for _, coord := range p.blocks {
				if !cube.Inside(coord[0], coord[1], coord[2]-1) {
					canFell = false
					continue
				}

				blockUnder := cube[coord[0]][coord[1]][coord[2]-1]
				if blockUnder != nil && blockUnder != p {
					canFell = false
					continue
				}
			}

			if canFell {
				fell = true
				for i, coord := range p.blocks {
					p.blocks[i] = [3]int{coord[0], coord[1], coord[2] - 1}
					cube[coord[0]][coord[1]][coord[2]] = nil
					cube[coord[0]][coord[1]][coord[2]-1] = p
				}
			}
		}

		if !fell {
			break
		}
	}

	supporting := sets.Set[*part]{}
	for _, p := range parts {
		thisSupporting := sets.Set[*part]{}

		for _, coord := range p.blocks {
			if !cube.Inside(coord[0], coord[1], coord[2]-1) {
				continue
			}
			blockUnder := cube[coord[0]][coord[1]][coord[2]-1]
			if blockUnder != nil && blockUnder != p {
				thisSupporting.Add(blockUnder)
			}
		}

		if len(thisSupporting) == 1 {
			for support := range thisSupporting {
				supporting.Add(support)
			}
		}
	}

	return len(parts) - len(supporting)
}
