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
	utils.Test(solve(test), 45000)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n\n")
	sums := []int{}
	for _, str := range s {
		sums = append(sums, math.Sum(sti.Stis(strings.Split(str, "\n"))))
	}

	sums = math.Order(sums, true)

	return math.Sum(sums[:3])
}
