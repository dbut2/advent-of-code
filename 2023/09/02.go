package main

import (
	"embed"
	"slices"

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
	h.Tester.Expect(1, 2)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	total := 0
	for _, line := range s {
		numbers := strings.Ints(line)
		slices.Reverse(numbers)
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

		nextValue := 0
		for _, layer := range layers {
			nextValue += layer[len(layer)-1]
		}
		total += nextValue
	}
	return total
}
