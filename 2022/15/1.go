package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve2(10))
	_ = t
	t.Expect(1, 26)
	fmt.Println(solve2(2000000)(input))
}

func solve2(y int) func(string) int {
	return func(s string) int {
		return solve(s, y)
	}
}

func solve(input string, y int) int {
	input = strings.ReplaceAll(input, ",", "")
	input = strings.ReplaceAll(input, ":", "")
	s := utils.ParseInput(input)

	var beacons []int
	var sensors []int
	var norange [][2]int

	for _, str := range s {
		split := strings.Split(str, " ")

		sx := sti.Sti(strings.Split(split[2], "=")[1])
		sy := sti.Sti(strings.Split(split[3], "=")[1])

		bx := sti.Sti(strings.Split(split[8], "=")[1])
		by := sti.Sti(strings.Split(split[9], "=")[1])

		if sy == y {
			sensors = append(sensors, sx)
		}

		if by == y {
			beacons = append(beacons, bx)
		}

		diff := math.Abs(sx-bx) + math.Abs(sy-by)

		start := sx - diff + math.Abs(sy-y)
		end := sx + diff - math.Abs(sy-y)

		if start <= end {
			norange = append(norange, [2]int{start, end})
		}
	}

	min := math.SmallestMap(norange, func(r [2]int) int {
		return r[0]
	})[0]

	max := math.LargestMap(norange, func(r [2]int) int {
		return r[1]
	})[1]

	fmt.Println(sensors)
	fmt.Println(beacons)

	count := 0
	for i := min; i <= max; i++ {
		counts := true
		for _, r := range norange {
			if i >= r[0] && i <= r[1] {
				counts = false
				break
			}
		}
		for _, b := range beacons {
			if b == i {
				counts = true
				break
			}
		}
		for _, s2 := range sensors {
			if s2 == i {
				counts = true
				break
			}
		}
		if !counts {
			count++
		}
	}

	return count
}

type Cell int

const (
	Empty Cell = iota
	Sensor
	Beacon
	NoBeacon
)
