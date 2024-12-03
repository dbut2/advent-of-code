package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Tester.Expect(1, 8)
	h.Run()
}

const (
	NS     uint8 = '|'
	EW     uint8 = '-'
	NE     uint8 = 'L'
	NW     uint8 = 'J'
	SW     uint8 = '7'
	SE     uint8 = 'F'
	Ground uint8 = '.'
	Start  uint8 = 'S'
)

func solve(input string) int {
	s := utils.ParseInput(input)

	grid := space.NewGridFromInput(s)

	start, _ := grid.Find(func(cell space.Cell, pipe uint8) bool {
		return pipe == Start
	})

	// determine start pipe shape
	{
		northCell := grid.Get(start.Move(space.North))
		southCell := grid.Get(start.Move(space.South))
		eastCell := grid.Get(start.Move(space.East))
		westCell := grid.Get(start.Move(space.West))

		north := northCell != nil && (*northCell == NS || *northCell == SW || *northCell == SE)
		south := southCell != nil && (*southCell == NS || *southCell == NE || *southCell == NW)
		east := eastCell != nil && (*eastCell == EW || *eastCell == NW || *eastCell == SW)
		west := westCell != nil && (*westCell == EW || *westCell == NE || *westCell == SE)

		switch {
		case north && south:
			grid.Set(start, NS)
		case east && west:
			grid.Set(start, EW)
		case north && east:
			grid.Set(start, NE)
		case north && west:
			grid.Set(start, NW)
		case south && west:
			grid.Set(start, SW)
		case south && east:
			grid.Set(start, SE)
		}
	}

	// create connections between pipes
	connectionDirections := map[uint8][2][2]int{
		NS: {space.North, space.South},
		EW: {space.West, space.East},
		NE: {space.North, space.East},
		NW: {space.North, space.West},
		SW: {space.South, space.West},
		SE: {space.South, space.East},
	}

	seen := sets.SetOf(start)
	loop := []space.Cell{start}
	queue := lists.NewQueue[space.Cell]()
	queue.Push(start.Move(connectionDirections[*grid.Get(start)][0]))

	for cell := range queue.Seq {
		if seen.Contains(cell) {
			continue
		}
		seen.Add(cell)

		pipe := grid.Get(cell)
		if pipe == nil {
			continue
		}

		loop = append(loop, cell)
		for _, direction := range connectionDirections[*grid.Get(cell)] {
			neighbor := cell.Move(direction)
			queue.Push(neighbor)
		}
	}

	return len(loop) / 2
}
