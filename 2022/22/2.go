package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/grid"
	"github.com/dbut2/advent-of-code/pkg/sti"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	//t := test.Register(tests, wrapSolve(4))
	//t.Expect(1, 6032)
	fmt.Println(solve(50, input))
}

func wrapSolve(gsize int) func(string) int {
	return func(input string) int {
		return solve(gsize, input)
	}
}

func solve(gsize int, input string) int {
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
			x, y, d = move(gsize, g, x, y, d, amt)
			d = turn(d, str)
		default:
			buffer = buffer + str
		}
	}

	amt := sti.Sti(buffer)
	buffer = ""
	x, y, d = move(gsize, g, x, y, d, amt)

	return 1000*(y+1) + 4*(x+1) + dirScore[d]
}

func move(gsize int, g grid.Grid[Cell], x, y int, d Direction, amt int) (int, int, Direction) {
	for i := 0; i < amt; i++ {
		c, nx, ny, nd := next(gsize, g, x, y, d)
		if c == Wall {
			return x, y, d
		}
		x, y, d = nx, ny, nd
	}
	return x, y, d
}

func next(gsize int, g grid.Grid[Cell], x, y int, d Direction) (Cell, int, int, Direction) {
	nx := x + d[0]
	ny := y + d[1]

	// if next cell not Unexist return it and new coords
	if g.Get(nx, ny) != Unexist {
		return g.Get(nx, ny), nx, ny, d
	}

	x, y, d = moveOverEdge(gsize, x, y, d)
	return g.Get(x, y), x, y, d
}

type Direction [2]int

var (
	Right Direction = [2]int{1, 0}
	Down  Direction = [2]int{0, 1}
	Left  Direction = [2]int{-1, 0}
	Up    Direction = [2]int{0, -1}
)

func (d Direction) opp() Direction {
	switch d {
	case Right:
		return Left
	case Down:
		return Up
	case Left:
		return Right
	case Up:
		return Down
	}
	panic("unknown direction")
}

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

func moveOverEdge(gsize int, x, y int, d Direction) (int, int, Direction) {
	// move around the net clockwise
	// a left turn is 2 sides that will get closed together
	// every side N edges away will be connected
	// keep track of how many corners until 1st left
	// continue that many times after left turn for matching side

	// ^cbf doing it algorithmically

	//  ..
	//  .
	// ..
	// .

	add, t := transform(gsize)

	add(1, 0, Up, 0, 3, Left)
	add(2, 0, Up, 0, 3, Down)
	add(2, 0, Right, 1, 2, Right)
	add(2, 0, Down, 1, 1, Right)
	add(1, 2, Down, 0, 3, Right)
	add(0, 2, Left, 1, 0, Left)
	add(0, 2, Up, 1, 1, Left)

	return t(x, y, d)

}

func transform(gsize int) (func(int, int, Direction, int, int, Direction), func(int, int, Direction) (int, int, Direction)) {
	type connection struct {
		in  [][2]int
		d1  Direction
		out [][2]int
		d2  Direction
	}
	var connections []connection

	add := func(x1, y1 int, d1 Direction, x2, y2 int, d2 Direction) {
		connections = append(connections, connection{
			in:  edgecoordsrange(gsize, x1, y1, d1),
			d1:  d1,
			out: reverserange(gsize, x2, y2, d2),
			d2:  d2.opp(),
		}, connection{
			in:  edgecoordsrange(gsize, x2, y2, d2),
			d1:  d2,
			out: reverserange(gsize, x1, y1, d1),
			d2:  d1.opp(),
		})
	}

	t := func(x, y int, d Direction) (int, int, Direction) {
		for _, c := range connections {
			if c.d1 != d {
				continue
			}
			for i, match := range c.in {
				if x == match[0] && y == match[1] {
					return c.out[i][0], c.out[i][1], c.d2
				}
			}
		}

		panic("could not find edge")
	}
	return add, t
}

func reverserange(gsize int, x, y int, d Direction) [][2]int {
	o := edgecoordsrange(gsize, x, y, d)
	var r [][2]int
	for i := len(o) - 1; i >= 0; i-- {
		r = append(r, o[i])
	}
	return r
}

func edgecoordsrange(gsize int, x, y int, d Direction) [][2]int {
	var r [][2]int
	switch d {
	case Right:
		for i := y * gsize; i < (y+1)*gsize; i++ {
			r = append(r, [2]int{(x+1)*gsize - 1, i})
		}
	case Down:
		for i := (x+1)*gsize - 1; i >= x*gsize; i-- {
			r = append(r, [2]int{i, (y+1)*gsize - 1})
		}
	case Left:
		for i := (y+1)*gsize - 1; i >= y*gsize; i-- {
			r = append(r, [2]int{x * gsize, i})
		}
	case Up:
		for i := x * gsize; i < (x+1)*gsize; i++ {
			r = append(r, [2]int{i, y * gsize})
		}
	}
	return r
}
