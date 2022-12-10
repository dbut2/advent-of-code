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
	t.Expected(1, 286)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")

	sx, sy := 0, 0
	wx, wy := 1, 10
	heading := 90

	for _, str := range s {
		line := strings.SplitN(str, "", 2)
		cmd := line[0]
		amt := sti.Sti(line[1])

		switch cmd {
		case "N":
			wx += amt
		case "E":
			wy += amt
		case "S":
			wx -= amt
		case "W":
			wy -= amt
		case "R":
			wx, wy = rotate(wx, wy, amt)
		case "L":
			wx, wy = rotate(wx, wy, -amt)
		case "F":
			switch heading {
			case 0:
				sx += wy * amt
				sy += -wx * amt
			case 90:
				sx += wx * amt
				sy += wy * amt
			case 180:
				sx += -wy * amt
				sy += wx * amt
			case 270:
				sx += -wx * amt
				sy += -wy * amt
			default:
				panic("angle not found")
			}
		}
	}

	return math.Abs(sx) + math.Abs(sy)
}

func rotate(x, y, rotation int) (int, int) {
	for i := 0; i < (rotation+360)%360; i += 90 {
		x, y = -y, x
	}
	return x, y
}
