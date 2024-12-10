package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/dbut2/advent-of-code/pkg/benchmark"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sets"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(2, 10)
	t.Expect(1, 58)
	fmt.Println(solve(input))
	benchmark.Run(func() {
		solve(input)
	}, benchmark.Time(time.Second*10))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	cubes := sets.Set[Cube]{}

	for _, str := range s {
		c := sti.Stis(strings.Split(str, ","))
		cube := Cube{x: c[0], y: c[1], z: c[2]}
		cubes.Add(cube)
	}

	allSides := sets.Set[Side]{}

	for cube := range cubes {
		for _, side := range cube.Sides() {
			if !cubes.Has(side.Facing()) {
				allSides.Add(side)
			}
		}
	}

	surfaces := generateSurfaces(allSides)
	max := 0
	for _, surface := range surfaces {
		max = math.Max(max, len(surface))
	}

	return max
}

// find all surfaces
func generateSurfaces(allSides sets.Set[Side]) []sets.Set[Side] {
	var sides sets.Set[Side]
	sides = allSides.Copy()

	var surfaces []sets.Set[Side]

	for len(sides) > 0 {
		var side Side
		for s := range sides {
			side = s
			break
		}

		surface := fillSurface(allSides, &sides, side)
		surfaces = append(surfaces, surface)
	}

	return surfaces
}

// find all sides on same surface as current side
func fillSurface(allSides sets.Set[Side], sides *sets.Set[Side], side Side) sets.Set[Side] {
	var surface sets.Set[Side]
	getNextAndRemove(allSides, sides, &surface, side)
	return surface
}

// find all the sides that continue on from given side, add to surface and remove from sides
func getNextAndRemove(allSides sets.Set[Side], sides *sets.Set[Side], surface *sets.Set[Side], side Side) {
	surface.Add(side)
	sides.Remove(side)

	for _, dirs := range possibleNext(side) {
		var next Side
		for i := 0; i < len(dirs); i++ {
			n := dirs[i]
			if allSides.Has(n) {
				next = n
				break
			}
		}

		if surface.Has(next) {
			continue
		}

		getNextAndRemove(allSides, sides, surface, next)
	}
}

// return ordered slice for each direction containing the next side to look for in surface
func possibleNext(side Side) [4][3]Side {
	var sides [4][3]Side
	for i, direction := range side.direction.Orthogonal() {
		c := direction.Opposite()
		sides[i] = [3]Side{
			{Cube: side.Cube.Move(direction).Move(side.direction), direction: c},
			{Cube: side.Cube.Move(direction), direction: side.direction},
			{Cube: side.Cube, direction: direction},
		}
	}
	return sides
}

type Cube struct {
	x, y, z int
}

func (c Cube) Sides() [6]Side {
	var sides [6]Side
	for i, direction := range Directions {
		sides[i] = Side{Cube: c, direction: direction}
	}
	return sides
}

func (c Cube) Move(d Direction) Cube {
	return Cube{
		x: c.x + d[0],
		y: c.y + d[1],
		z: c.z + d[2],
	}
}

type Side struct {
	Cube
	direction Direction
}

func (s Side) Facing() Cube {
	return s.Cube.Move(s.direction)
}

func (d Direction) Orthogonal() [4]Direction {
	var o [4]Direction
	i := 0
	for _, d2 := range Directions {
		if d2 != d && d2 != d.Opposite() {
			o[i] = d2
			i++
		}
	}
	return o
}

type Direction [3]int

var (
	Right    Direction = [3]int{1, 0, 0}
	Up       Direction = [3]int{0, 1, 0}
	Forward  Direction = [3]int{0, 0, 1}
	Left     Direction = [3]int{-1, 0, 0}
	Down     Direction = [3]int{0, -1, 0}
	Backward Direction = [3]int{0, 0, -1}
)

var Directions = [6]Direction{Right, Up, Forward, Left, Down, Backward}

func (d Direction) Opposite() Direction {
	return Direction{0 - d[0], 0 - d[1], 0 - d[2]}
}
