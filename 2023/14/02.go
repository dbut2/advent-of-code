package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Tester.Expect(1, 64)
	h.Run()
}

type rock uint8

const (
	None  rock = '.'
	Round rock = 'O'
	Cube  rock = '#'
)

func solve(input string) int {
	s := utils.ParseInput(input)

	seen := map[string]int{}

	var cycleStart, cycleEnd int
	for n := 1; ; n++ {
		cycle(s)

		if lastSeen, ok := seen[strings.Join(s, "\n")]; ok {
			cycleStart = lastSeen
			cycleEnd = n
			break
		}
		seen[strings.Join(s, "\n")] = n
	}

	remaining := (1e9 - cycleStart) % (cycleEnd - cycleStart)
	for n := 0; n < remaining; n++ {
		cycle(s)
	}

	total := 0
	for j := range s {
		for i := range s[j] {
			if s[j][i] == uint8(Round) {
				total += len(s) - j
			}
		}
	}
	return total
}

func cycle(s []string) {
	x2 := len(s[0])
	y2 := len(s)

	for i := 0; i < x2; i++ {
		topFree := 0
		for j := 0; j < y2; j++ {
			switch s[j][i] {
			case uint8(None):
				continue
			case uint8(Cube):
				topFree = j + 1
			case uint8(Round):
				s[j] = s[j][:i] + string(None) + s[j][i+1:]
				s[topFree] = s[topFree][:i] + string(Round) + s[topFree][i+1:]
				topFree++
			}
		}
	}

	for j := 0; j < y2; j++ {
		leftFree := 0
		for i := 0; i < x2; i++ {
			switch s[j][i] {
			case uint8(None):
				continue
			case uint8(Cube):
				leftFree = i + 1
			case uint8(Round):
				s[j] = s[j][:i] + string(None) + s[j][i+1:]
				s[j] = s[j][:leftFree] + string(Round) + s[j][leftFree+1:]
				leftFree++
			}
		}
	}

	for i := x2 - 1; i >= 0; i-- {
		bottomFree := y2 - 1
		for j := y2 - 1; j >= 0; j-- {
			switch s[j][i] {
			case uint8(None):
				continue
			case uint8(Cube):
				bottomFree = j - 1
			case uint8(Round):
				s[j] = s[j][:i] + string(None) + s[j][i+1:]
				s[bottomFree] = s[bottomFree][:i] + string(Round) + s[bottomFree][i+1:]
				bottomFree--
			}
		}
	}

	for j := y2 - 1; j >= 0; j-- {
		rightFree := x2 - 1
		for i := x2 - 1; i >= 0; i-- {
			switch s[j][i] {
			case uint8(None):
				continue
			case uint8(Cube):
				rightFree = i - 1
			case uint8(Round):
				s[j] = s[j][:i] + string(None) + s[j][i+1:]
				s[j] = s[j][:rightFree] + string(Round) + s[j][rightFree+1:]
				rightFree--
			}
		}
	}
}
