package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), 19208)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")
	adapters := math.Order(sti.Stis(s), false)

	waysTo := map[int]int{0: 1}

	for _, adapter := range adapters {
		for i := 0; i <= 3; i++ {
			if ways, ok := waysTo[adapter-i]; ok {
				waysTo[adapter] += ways
			}
		}
	}

	return waysTo[math.Largest(adapters)]
}
