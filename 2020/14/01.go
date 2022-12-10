package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expected(1, 165)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")

	var maskOn int
	var maskOff int
	mem := make(map[int]int)
	for _, str := range s {

		if str[:4] == "mask" {
			maskOn, maskOff = parseMasks(str[7:])
		}

		if str[:3] == "mem" {
			line := strings.Split(str[3:], " = ")
			addr := sti.Sti(strings.Trim(line[0], "[]"))
			val := sti.Sti(line[1])

			mem[addr] = val&maskOff | maskOn
		}
	}

	return math.SumMap(mem)
}

func parseMasks(mask string) (int, int) {
	maskOn := 0
	maskOff := math.Pow(2, 36) - 1

	maskPlaces := strings.Split(mask, "")
	for i := 0; i < len(maskPlaces); i++ {
		m := maskPlaces[len(maskPlaces)-1-i]

		switch m {
		case "1":
			maskOn += math.Pow(2, i)
		case "0":
			maskOff -= math.Pow(2, i)
		}
	}

	return maskOn, maskOff
}
