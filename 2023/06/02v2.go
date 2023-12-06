package main

import (
	"embed"
	"fmt"
	"math"

	"github.com/dbut2/advent-of-code/pkg/buffers"
	"github.com/dbut2/advent-of-code/pkg/chars"
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

	buffer := buffers.Number(0)
	for _, char := range s[0] {
		if chars.IsNum(char) {
			buffer.Add(chars.NumVal(char))
		}
	}
	time := float64(buffer.Clear())
	for _, char := range s[1] {
		if chars.IsNum(char) {
			buffer.Add(chars.NumVal(char))
		}
	}
	distance := float64(buffer.Clear())

	discriminant := time*time - 4*distance

	root1 := (time + math.Sqrt(discriminant)) / 2
	root2 := (time - math.Sqrt(discriminant)) / 2

	perms := int(math.Floor(root1)-math.Ceil(root2)) + 1

	return perms
}
