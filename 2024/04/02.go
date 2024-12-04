package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/space"
)

func solve(input space.Grid[byte]) int {
	count := 0

	for coords, cell := range input.Cells() {
		if *cell != 'A' {
			continue
		}

		dirs := space.Diagonals
		for range 4 {
			a, b, c, d := input.Get(coords.Move(dirs[0])), input.Get(coords.Move(dirs[1])), input.Get(coords.Move(dirs[2])), input.Get(coords.Move(dirs[3]))
			if a == nil || b == nil || c == nil || d == nil {
				continue
			}

			if *a == 'M' && *b == 'M' && *c == 'S' && *d == 'S' {
				count++
			}

			dirs = rotate(dirs)
		}
	}

	return count
}

func rotate(dirs []space.Direction) []space.Direction {
	return append(dirs[1:], dirs[0])
}

func main() {
	h := harness.New(solve, inputs)
	h.Run()
}

//go:embed *.txt
var inputs embed.FS
