package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/grid"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 3068)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	var directions []int

	for _, d := range strings.Split(s[0], "") {
		switch d {
		case "<":
			directions = append(directions, left)
		case ">":
			directions = append(directions, right)
		default:
			panic(d)
		}
	}

	g := grid.Grid[bool]{}

	m := -1
	for n := 0; n < 2022; n++ {
		rock := Rocks[n%len(Rocks)]
		i, j := findTop(g)+4, 2

		for {
			m++
			direction := directions[m%len(directions)]

			// can be blown
			canBeBlown := true
			for x, row := range rock {
				x += i
				for y, cell := range row {
					y += j

					if !cell {
						continue
					}
					newY := y + direction

					if newY < 0 || newY >= 7 {
						canBeBlown = false
						break
					}

					if g.Get(x, newY) {
						canBeBlown = false
						break
					}
				}
			}

			if canBeBlown {
				j += direction
			}

			// can move down
			canMoveDown := true
			for x, row := range rock {
				x += i
				for y, cell := range row {
					y += j

					if !cell {
						continue
					}

					newX := x - 1

					if newX < 0 {
						canMoveDown = false
						break
					}

					if g.Get(newX, y) {
						canMoveDown = false
						break
					}
				}
			}

			if canMoveDown {
				i -= 1
			}

			if !canMoveDown {
				break
			}
		}

		//for a := len(g) - 1; a >= 0; a-- {
		//	for b := 0; b < 7; b++ {
		//		if g.Get(a, b) {
		//			fmt.Print("#")
		//		} else {
		//			fmt.Print(".")
		//		}
		//	}
		//	fmt.Println()
		//}
		//fmt.Println(i, j)
		//fmt.Println()

		for x, row := range rock {
			for y, cell := range row {
				if cell {
					g.Set(x+i, y+j, true)
				}
			}
		}
	}

	return findTop(g) + 1
}

// 0   1   2
// >>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>

func findTop(g grid.Grid[bool]) int {
	for i := 0; ; i++ {
		rowHas := false
		for j := 0; j < 7; j++ {
			if g.Get(i, j) {
				rowHas = true
			}
		}
		if !rowHas {
			return i - 1
		}
	}
}

const (
	left int = iota - 1
	_
	right
)

type Rock [][]bool

var Rocks = []Rock{
	{
		{true, true, true, true},
	},
	{
		{false, true, false},
		{true, true, true},
		{false, true, false},
	},
	{
		{true, true, true},
		{false, false, true},
		{false, false, true},
	},
	{
		{true},
		{true},
		{true},
		{true},
	},
	{
		{true, true},
		{true, true},
	},
}
