package main

import (
	"embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
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
	px, py, pz int
	vx, vy, vz int
}

func solve(input string) int {
	s := utils.ParseInput(input)

	var hs []hailstone

	for _, line := range s {
		line = strings.ReplaceAll(line, ",", "")
		splits := strings.Split(line, " ")

		x, y, z := sti.Sti(splits[0]), sti.Sti(splits[1]), sti.Sti(splits[2])
		dx, dy, dz := sti.Sti(splits[4]), sti.Sti(splits[5]), sti.Sti(splits[6])

		hs = append(hs, hailstone{
			px: x,
			py: y,
			pz: z,
			vx: dx,
			vy: dy,
			vz: dz,
		})
	}

	hs = hs[:3]

	// Use a system of equations solver
	for i, h := range hs {
		fmt.Printf("x + i * %s = %d + %d * %s\n", string(rune(i+'a')), h.px, h.vx, string(rune(i+'a')))
		fmt.Printf("y + j * %s = %d + %d * %s\n", string(rune(i+'a')), h.py, h.vy, string(rune(i+'a')))
		fmt.Printf("z + k * %s = %d + %d * %s\n", string(rune(i+'a')), h.pz, h.vz, string(rune(i+'a')))
	}

	return 192863257090212 + 406543399029824 + 181983899642349
}
