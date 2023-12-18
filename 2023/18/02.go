package main

import (
	"embed"
	"slices"
	"strings"

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
	h.Expect(1, 952408144115)
	h.Solve()
}

type direction string

const (
	up    direction = "U"
	down  direction = "D"
	left  direction = "L"
	right direction = "R"
)

type action struct {
	direction direction
	amount    int
	colour    string
}

func hexToDec(hex string) int {
	dec := 0
	for _, char := range hex {
		dec *= 16
		if char >= '0' && char <= '9' {
			dec += int(char - '0')
		}
		if char >= 'a' && char <= 'f' {
			dec += int(char-'a') + 10
		}
	}
	return dec
}

func solve(input string) int {
	s := utils.ParseInput(input)

	actions := []action{}

	for j, line := range s {
		_, _ = j, line

		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		line = strings.ReplaceAll(line, "#", "")

		parts := strings.Split(line, " ")

		a := action{
			amount: hexToDec(parts[2][:5]),
		}

		switch parts[2][5] {
		case '0':
			a.direction = right
		case '1':
			a.direction = down
		case '2':
			a.direction = left
		case '3':
			a.direction = up
		}

		actions = append(actions, a)
	}

	coords := [2]int{0, 0}

	lines := sets.Set[[2][2]int]{}

	for _, a := range actions {
		switch a.direction {
		case up:
			x1, y1 := coords[0], coords[1]
			coords[0] -= a.amount
			x2, y2 := coords[0], coords[1]
			lines.Add([2][2]int{{x1, y1}, {x2, y2}})
		case down:
			x1, y1 := coords[0], coords[1]
			coords[0] += a.amount
			x2, y2 := coords[0], coords[1]
			lines.Add([2][2]int{{x1, y1}, {x2, y2}})
		case left:
			x1, y1 := coords[0], coords[1]
			coords[1] -= a.amount
			x2, y2 := coords[0], coords[1]
			lines.Add([2][2]int{{x1, y1}, {x2, y2}})
		case right:
			x1, y1 := coords[0], coords[1]
			coords[1] += a.amount
			x2, y2 := coords[0], coords[1]
			lines.Add([2][2]int{{x1, y1}, {x2, y2}})
		}
	}

	xPoints := []int{}
	yPoints := []int{}

	for line := range lines {
		xPoints = append(xPoints, line[0][0])
		xPoints = append(xPoints, line[0][0]+1)
		xPoints = append(xPoints, line[1][0])
		xPoints = append(xPoints, line[1][0]+1)
		yPoints = append(yPoints, line[0][1])
		yPoints = append(yPoints, line[0][1]+1)
		yPoints = append(yPoints, line[1][1])
		yPoints = append(yPoints, line[1][1]+1)
	}

	xps := sets.SetFrom(xPoints)
	yps := sets.SetFrom(yPoints)

	xPoints = xps.Slice()
	yPoints = yps.Slice()

	slices.Sort(xPoints)
	slices.Sort(yPoints)

	ySizes := []int{}
	xSizes := []int{}

	for i := 1; i < len(xPoints); i++ {
		xSizes = append(xSizes, xPoints[i]-xPoints[i-1])
	}

	for i := 1; i < len(yPoints); i++ {
		ySizes = append(ySizes, yPoints[i]-yPoints[i-1])
	}

	xPoints = xPoints[:len(xPoints)-1]
	yPoints = yPoints[:len(yPoints)-1]

	mapToBig := func(min [2]int) [2]int {
		return [2]int{xPoints[min[0]], yPoints[min[1]]}
	}

	outside := sets.Set[[2]int]{}

	check := lists.Queue[[2]int]{}

	for x := 0; x < len(xPoints); x++ {
		check.Push([2]int{x, 0})
		check.Push([2]int{x, len(yPoints) - 1})
	}

	for y := 0; y < len(yPoints); y++ {
		check.Push([2]int{0, y})
		check.Push([2]int{len(xPoints) - 1, y})
	}

	for len(check) > 0 {
		item := check.Pop()

		if item[0] < 0 || item[1] < 0 || item[0] >= len(xPoints) || item[1] >= len(yPoints) {
			continue
		}

		if outside.Has(item) {
			continue
		}

		if inLine(lines.Slice(), mapToBig(item)) {
			continue
		}

		outside.Add(item)

		for _, delta := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			check.Push([2]int{item[0] + delta[0], item[1] + delta[1]})
		}
	}

	total := 0
	for y := 0; y < len(yPoints); y++ {
		for x := 0; x < len(xPoints); x++ {
			if !outside.Has([2]int{x, y}) {
				total += (xSizes[x]) * (ySizes[y])
			}
		}
	}

	return total
}

func inLine(lines [][2][2]int, coord [2]int) bool {
	x, y := coord[0], coord[1]
	for _, line := range lines {
		x1, y1 := line[0][0], line[0][1]
		x2, y2 := line[1][0], line[1][1]

		x1, x2 = min(x1, x2), max(x1, x2)
		y1, y2 = min(y1, y2), max(y1, y2)

		if x1 == x2 {
			if x == x1 && y >= y1 && y <= y2 {
				return true
			}
		}

		if y1 == y2 {
			if y == y1 && x >= x1 && x <= x2 {
				return true
			}
		}
	}
	return false
}
