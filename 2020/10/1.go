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
	utils.Test(solve(test), 220)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")

	a := math.Order(sti.Stis(s), false)

	o := 1
	t := 1

	for i := 1; i < len(a); i++ {
		if a[i]-a[i-1] == 1 {
			o++
		}
		if a[i]-a[i-1] == 3 {
			t++
		}
	}

	return o * t
}
