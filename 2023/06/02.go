package main

import (
	"embed"
	"math"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	strings2 "github.com/dbut2/advent-of-code/pkg/strings"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 71503)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	s[0] = strings.ReplaceAll(s[0], " ", "")
	time := float64(strings2.Ints(s[0])[0])
	s[1] = strings.ReplaceAll(s[1], " ", "")
	distance := float64(strings2.Ints(s[1])[0])

	discriminant := time*time - 4*distance

	root1 := (time + math.Sqrt(discriminant)) / 2
	root2 := (time - math.Sqrt(discriminant)) / 2

	perms := int(math.Floor(root1)-math.Ceil(root2)) + 1

	return perms
}
