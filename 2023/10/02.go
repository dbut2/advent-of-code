package main

import (
	"embed"
	"github.com/dbut2/advent-of-code/pkg/grid"
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(2, 10)
	h.Solve()
}

type Pipe struct {
	shape       Shape
	x, y        int
	connections [2]*Pipe
}

type Shape rune

const (
	NS     Shape = '|'
	EW     Shape = '-'
	NE     Shape = 'L'
	NW     Shape = 'J'
	SW     Shape = '7'
	SE     Shape = 'F'
	Ground Shape = '.'
	Start  Shape = 'S'
)

func solve(input string) int {
	s := utils.ParseInput(input)

	g := grid.Grid[Pipe]{}

	for i, line := range s {
		for j, char := range line {
			p := Pipe{
				shape: Shape(char),
				x:     i,
				y:     j,
			}
			g.Set(i, j, p)
		}
	}

	start := g.Find(func(pipe Pipe) bool {
		return pipe.shape == Start
	})

	// determine start tile shape
	{
		northCell := g.Get(start.x-1, start.y)
		southCell := g.Get(start.x+1, start.y)
		eastCell := g.Get(start.x, start.y+1)
		westCell := g.Get(start.x, start.y-1)

		north := northCell != nil && (northCell.shape == NS || northCell.shape == SW || northCell.shape == SE)
		south := southCell != nil && (southCell.shape == NS || southCell.shape == NE || northCell.shape == NW)
		east := eastCell != nil && (eastCell.shape == EW || eastCell.shape == NW || eastCell.shape == SW)
		west := westCell != nil && (westCell.shape == EW || westCell.shape == NE || westCell.shape == SE)

		switch {
		case north && south:
			start.shape = NS
		case east && west:
			start.shape = EW
		case north && east:
			start.shape = NE
		case north && west:
			start.shape = NW
		case south && west:
			start.shape = SW
		case south && east:
			start.shape = SE
		}
	}

	// create connections between pipes
	connectionOffsets := map[Shape][2][2]int{
		NS: {{-1, 0}, {1, 0}},
		EW: {{0, 1}, {0, -1}},
		NE: {{-1, 0}, {0, 1}},
		NW: {{-1, 0}, {0, -1}},
		SW: {{1, 0}, {0, -1}},
		SE: {{1, 0}, {0, 1}},
	}
	for _, pipe := range g {
		x1 := connectionOffsets[pipe.shape][0][0]
		y1 := connectionOffsets[pipe.shape][0][1]
		x2 := connectionOffsets[pipe.shape][1][0]
		y2 := connectionOffsets[pipe.shape][1][1]

		pipe.connections = [2]*Pipe{
			g.Get(pipe.x+x1, pipe.y+y1),
			g.Get(pipe.x+x2, pipe.y+y2),
		}
	}

	// find all pipes in loop
	loop := sets.Set[*Pipe]{}
	queue := lists.Queue[*Pipe]{}
	queue.Push(start.connections[0], start.connections[1])
	for len(queue) > 0 {
		pipe := queue.Pop()
		if loop.Has(pipe) {
			continue
		}
		loop.Add(pipe)
		queue.Push(pipe.connections[0], pipe.connections[1])
	}

	// determine if each tile is enclosed by loop
	// use the count of crossed loop pipes between tile and border, we only need to check a single direction
	// an odd count means the tile is enclosed
	// we only really check the right half of horizontal pipes so that L and 7, and J and F don't cancel each other out
	// eg below we can see there is only effectively 1 horizontal crossing
	// .F-
	// .|.
	// -J.
	// by only counting pipes with a right segment, F in this case, we won't accidentally double count single crossings

	enclosed := 0
	for coords, tile := range g {
		if loop.Has(tile) {
			continue
		}

		pipesCrossed := 0
		for x := coords[0]; x >= 0; x-- {
			checkPipe := g.Get(x, coords[1])
			if !loop.Has(checkPipe) {
				continue
			}

			if checkPipe.shape == EW || checkPipe.shape == NE || checkPipe.shape == SE {
				pipesCrossed++
			}
		}

		if pipesCrossed%2 == 1 {
			enclosed++
		}
	}

	return enclosed
}
