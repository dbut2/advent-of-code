package main

import (
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/sti"
)

type grid = space.Grid[byte]

func solve(input [2][]string) int {
	total := 0

	keys := []string{}

	values := make(map[string]int)

	for _, line := range input[0] {
		parts := strings.Split(line, ": ")
		values[parts[0]] = sti.Int(parts[1])
		keys = append(keys, parts[0])
	}

	type calc struct {
		a, b string
		op   string
	}

	calcs := map[string]calc{}
	for _, line := range input[1] {
		parts := strings.Split(line, " ")
		calcs[parts[4]] = calc{a: parts[0], b: parts[2], op: parts[1]}
		keys = append(keys, parts[4])
	}

	var calculate func(key string) int
	calculate = func(key string) int {
		if v, ok := values[key]; ok {
			return v
		}

		c := calcs[key]
		a := calculate(c.a)
		b := calculate(c.b)

		var out int
		switch c.op {
		case "AND":
			out = a & b
		case "OR":
			out = a | b
		case "XOR":
			out = a ^ b
		default:
			panic("unknown op " + c.op)
		}

		values[key] = out
		return out
	}

	for _, key := range keys {
		if strings.HasPrefix(key, "z") {
			n := strings.TrimPrefix(key, "z")
			total += calculate(key) * math.Pow(2, sti.Int(n))
		}
	}

	return total
}

func main() {
	h := harness.New(solve)
	h.Expect(1, 2024)
	h.Run()
}
