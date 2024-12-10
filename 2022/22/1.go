package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/grid"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 6032)
	fmt.Println(solve(input))
}

func solve(input string) int {
	parts := strings.Split(input, "\n\n")
	s := strings.Split(parts[0], "\n")
	instructions := parts[1]

	g := grid.Grid[Cell]{}

	for y, line := range s {
		for x, str := range strings.Split(line, "") {
			c := strCellMap[str]
			g.Set(x, y, c)
		}
	}

	x, y := 0, 0
	for g.Get(x, y) == Unexist {
		x += 1
	}
	d := Right

	buffer := ""
	for _, str := range strings.Split(instructions, "") {
		switch str {
		case "L", "R":
			amt := sti.Sti(buffer)
			buffer = ""
			x, y = move(g, x, y, d, amt)
			d = turn(d, str)
		default:
			buffer = buffer + str
		}
	}

	amt := sti.Sti(buffer)
	buffer = ""
	x, y = move(g, x, y, d, amt)

	return 1000*(y+1) + 4*(x+1) + dirScore[d]
}

func move(g grid.Grid[Cell], x, y int, d Direction, amt int) (int, int) {
	for i := 0; i < amt; i++ {
		c, nx, ny := next(g, x, y, d)
		if c == Wall {
			return x, y
		}
		x, y = nx, ny
	}
	return x, y
}

func next(g grid.Grid[Cell], x, y int, d Direction) (Cell, int, int) {
	x += d[0]
	y += d[1]

	// if next cell not Unexist return it and new coords
	if g.Get(x, y) != Unexist {
		return g.Get(x, y), x, y
	}

	// else we hit edge, move backwards 200
	x -= 200 * d[0]
	y -= 200 * d[1]

	// move forwards while the cell is Unexist
	for g.Get(x, y) == Unexist {
		x += d[0]
		y += d[1]
	}

	return g.Get(x, y), x, y
}

type Direction [2]int

var (
	Right Direction = [2]int{1, 0}
	Down  Direction = [2]int{0, 1}
	Left  Direction = [2]int{-1, 0}
	Up    Direction = [2]int{0, -1}
)

var dirScore = map[Direction]int{
	Right: 0,
	Down:  1,
	Left:  2,
	Up:    3,
}

func turn(current Direction, turn string) Direction {
	switch turn {
	case "L":
		switch current {
		case Right:
			return Up
		case Up:
			return Left
		case Left:
			return Down
		case Down:
			return Right
		}
	case "R":
		switch current {
		case Right:
			return Down
		case Down:
			return Left
		case Left:
			return Up
		case Up:
			return Right
		}
	}
	panic("unknown turn")
}

type Cell int

const (
	Unexist Cell = iota
	Open
	Wall
)

var strCellMap = map[string]Cell{
	" ": Unexist,
	".": Open,
	"#": Wall,
}
