package main

import (
	"embed"
	"math"

	"github.com/dbut2/advent-of-code/pkg/chars"
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(1, 94)
	h.Expect(2, 71)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	heatLoss := space.NewGrid[int](len(s[0]), len(s))
	minsA := space.NewGrid[int](len(s[0]), len(s))
	minsB := space.NewGrid[int](len(s[0]), len(s))

	for j, line := range s {
		_, _ = j, line

		for i, char := range line {
			heatLoss[i][j] = chars.NumVal(char)
			minsA[i][j] = math.MaxInt
			minsB[i][j] = math.MaxInt
		}
	}

	type update struct {
		x, y         int
		newMin       int
		onHorizontal bool
	}

	queue := lists.Queue[update]{}

	queue.Push(update{0, 0, 0, true})
	queue.Push(update{0, 0, 0, false})

	for len(queue) > 0 {
		item := queue.Pop()

		mins := minsA
		if item.onHorizontal {
			mins = minsB
		}

		if item.newMin >= mins[item.x][item.y] {
			continue
		}

		mins[item.x][item.y] = item.newMin

		switch item.onHorizontal {
		case true:
			runningTotal := item.newMin
			for i := 1; i <= 10; i++ {
				x, y := item.x+i, item.y
				if heatLoss.Inside(x, y) {
					runningTotal += heatLoss[x][y]
					if i >= 4 {
						queue.Push(update{x, y, runningTotal, false})
					}
				} else {
					break
				}
			}

			runningTotal = item.newMin
			for i := 1; i <= 10; i++ {
				x, y := item.x-i, item.y
				if heatLoss.Inside(x, y) {
					runningTotal += heatLoss[x][y]
					if i >= 4 {
						queue.Push(update{x, y, runningTotal, false})
					}
				} else {
					break
				}
			}
		case false:
			runningTotal := item.newMin
			for i := 1; i <= 10; i++ {
				x, y := item.x, item.y+i
				if heatLoss.Inside(x, y) {
					runningTotal += heatLoss[x][y]
					if i >= 4 {
						queue.Push(update{x, y, runningTotal, true})
					}
				} else {
					break
				}
			}

			runningTotal = item.newMin
			for i := 1; i <= 10; i++ {
				x, y := item.x, item.y-i
				if heatLoss.Inside(x, y) {
					runningTotal += heatLoss[x][y]
					if i >= 4 {
						queue.Push(update{x, y, runningTotal, true})
					}
				} else {
					break
				}
			}
		}
	}

	endNodeA := minsA[len(heatLoss)-1][len(heatLoss[0])-1]
	endNodeB := minsB[len(heatLoss)-1][len(heatLoss[0])-1]

	return min(endNodeA, endNodeB)
}
