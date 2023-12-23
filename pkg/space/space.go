package space

type Cell [2]int

func (c Cell) Move(d Direction) Cell {
	return Cell{c[0] + d[0], c[1] + d[1]}
}

func (c Cell) DirectionTo(b Cell) Direction {
	return Direction{b[0] - c[0], b[1] - c[1]}
}

type Direction [2]int

func (d Direction) Add(b Direction) Direction {
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
