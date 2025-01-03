package space

import (
	"github.com/dbut2/advent-of-code/pkg/math"
)

type Cell [2]int

type Cells [][2]int

func (c Cell) Move(d Direction) Cell {
	return Cell{c[0] + d[0], c[1] + d[1]}
}

func (c Cell) DirectionTo(b Cell) Direction {
	return Direction{b[0] - c[0], b[1] - c[1]}
}

func Manhattan(a, b Cell) int {
	return math.Abs(a[0]-b[0]) + math.Abs(a[1]-b[1])
}

func Coord(cell Cell) [2]int {
	return [2]int{cell[0], cell[1]}
}

func (c Cells) Coords() [][2]int {
	cs := make([][2]int, 0, len(c))
	for _, cell := range c {
		cs = append(cs, Coord(cell))
	}
	return cs
}

func Coords(cells []Cell) [][2]int {
	coord := make([][2]int, 0, len(cells))
	for _, cell := range cells {
		coord = append(coord, Coord(cell))
	}
	return coord
}

type Direction [2]int

func (d Direction) Rotate() Direction {
	return Direction{-d[1], d[0]}
}

func (d Direction) Add(b Direction) Direction {
	return Direction{d[0] + b[0], d[1] + b[1]}
}

func (d Direction) Multiply(n int) Direction {
	return Direction{d[0] * n, d[1] * n}
}

var (
	North = Direction{0, -1}
	South = Direction{0, 1}
	East  = Direction{1, 0}
	West  = Direction{-1, 0}
	Up    = North
	Down  = South
	Left  = West
	Right = East
)

var Directions = []Direction{Up, Down, Left, Right}
var Diagonals = []Direction{Up.Add(Left), Up.Add(Right), Down.Add(Right), Down.Add(Left)}
