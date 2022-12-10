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
	t.Expected(2, 208)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")

	var mask string
	mem := make(map[int]int)
	for _, str := range s {
		if str[:4] == "mask" {
			mask = str[7:]
		}

		if str[:3] == "mem" {
			line := strings.Split(str[3:], " = ")
			addr := sti.Sti(strings.Trim(line[0], "[]"))
			val := sti.Sti(line[1])

			addrs := getAddrs(addr, mask)
			for _, a := range addrs {
				mem[a] = val
			}
		}
	}

	return math.SumMap(mem)
}

func getAddrs(addr int, mask string) []int {
	addrs := []int{addr}

	maskPlaces := strings.Split(mask, "")
	for i := 0; i < len(maskPlaces); i++ {
		m := maskPlaces[len(maskPlaces)-1-i]
		switch m {
		case "1":
			var newAddrs []int
			for _, a := range addrs {
				newAddrs = append(newAddrs, a|math.Pow(2, i))
			}
			addrs = newAddrs
		case "X":
			var newAddrs []int
			for _, a := range addrs {
				first := a | math.Pow(2, i)
				second := first - math.Pow(2, i)
				newAddrs = append(newAddrs, first, second)
			}
			addrs = newAddrs
		}
	}

	return addrs
}
