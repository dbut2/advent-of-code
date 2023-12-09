package main

import (
	"embed"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/strings"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Expect(1, 2)
	h.Solve()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	total := 0

	for _, line := range s {
		numbers := strings.Ints(line)
		layers := [][]int{numbers}

		for {
			bottomLayer := layers[len(layers)-1]

			all0 := true
			for _, val := range bottomLayer {
				if val != 0 {
					all0 = false
					break
				}
			}
			if all0 {
				break
			}

			nextLayer := make([]int, len(bottomLayer)-1)
			for i := range nextLayer {
				nextLayer[i] = bottomLayer[i+1] - bottomLayer[i]
			}

			layers = append(layers, nextLayer)
		}

		prevValue := 0
		for i := len(layers) - 1; i >= 0; i-- {
			prevValue = layers[i][0] - prevValue
		}
		total += prevValue
	}
	return total
}
