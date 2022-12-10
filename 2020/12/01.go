package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expected(1, 25)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")

	x, y := 0, 0
	heading := 90

	for _, str := range s {
		line := strings.SplitN(str, "", 2)
		cmd := line[0]
		amt := sti.Sti(line[1])

		switch cmd {
		case "N":
			x += amt
		case "E":
			y += amt
		case "S":
			x -= amt
		case "W":
			y -= amt
		case "R":
			heading = rotate(heading, amt)
		case "L":
			heading = rotate(heading, -amt)
		case "F":
			switch heading {
			case 0:
				x += amt
			case 90:
				y += amt
			case 180:
				x -= amt
			case 270:
				y -= amt
			default:
				panic("angle not found")
			}
		}
	}

	return math.Abs(x) + math.Abs(y)
}

func rotate(heading int, rotation int) int {
	heading += rotation
	for heading < 0 {
		heading += 360
	}
	return heading % 360
}
