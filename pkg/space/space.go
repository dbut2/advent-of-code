package space

type Cell [2]int

type Cells [][2]int

func (c Cell) Move(d Direction) Cell {
	return Cell{c[0] + d[0], c[1] + d[1]}
}

func (c Cell) DirectionTo(b Cell) Direction {
	return Direction{b[0] - c[0], b[1] - c[1]}
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

func (d *Direction) Add(b Direction) Direction {
	return Direction{d[0] + b[0], d[1] + b[1]}
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
