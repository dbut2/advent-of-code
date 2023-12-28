package main

import (
	"embed"
	"fmt"
	"math"
	"strings"

	strings2 "github.com/dbut2/advent-of-code/pkg/strings"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 71503)
	fmt.Println(solve(input))
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
