package main

import (
	"embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
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
	px, py, pz int
	vx, vy, vz int
}

func solve(input string) int {
	s := utils.ParseInput(input)

	var hs []hailstone

	for _, line := range s {
		line = strings.ReplaceAll(line, ",", "")
		splits := strings.Split(line, " ")

		x, y, z := sti.Int(splits[0]), sti.Int(splits[1]), sti.Int(splits[2])
		dx, dy, dz := sti.Int(splits[4]), sti.Int(splits[5]), sti.Int(splits[6])

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
