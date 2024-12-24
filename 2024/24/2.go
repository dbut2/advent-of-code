package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
)

func solve(input [2][]string) string {
	type calc struct {
		name string
		a, b string
		op   string
	}
	var calcs []calc
	find := func(a, b string, op string) (calc, bool) {
		if a > b {
			a, b = b, a
		}
		return lists.Find(calcs, func(calc calc) bool {
			return calc.op == op && calc.a == a && calc.b == b
		})
	}

	for _, gate := range input[1] {
		parts := strings.Split(gate, " ")
		a, b := parts[0], parts[2]
		if b < a {
			a, b = b, a
		}
		calcs = append(calcs, calc{
			name: parts[4],
			a:    a,
			b:    b,
			op:   parts[1],
		})
	}

	var swapped []string

	var carry calc
	for i := range 45 {
		n := fmt.Sprintf("%02d", i)
		xor, _ := find("x"+n, "y"+n, "XOR")
		and, _ := find("x"+n, "y"+n, "AND")

		if carry.name == "" {
			carry = and
			continue
		}

		gate1, ok := find(carry.name, xor.name, "AND")
		if !ok {
			xor, and = and, xor
			swapped = append(swapped, and.name, xor.name)
			gate1, _ = find(carry.name, xor.name, "AND")
		}

		z, ok := find(carry.name, xor.name, "XOR")
		if strings.HasPrefix(xor.name, "z") {
			xor, z = z, xor
			swapped = append(swapped, xor.name, z.name)
		}
		if strings.HasPrefix(and.name, "z") {
			and, z = z, and
			swapped = append(swapped, and.name, z.name)
		}
		if strings.HasPrefix(gate1.name, "z") {
			gate1, z = z, gate1
			swapped = append(swapped, gate1.name, z.name)
		}

		carry, _ = find(gate1.name, and.name, "OR")
		if strings.HasPrefix(carry.name, "z") {
			carry, z = z, carry
			swapped = append(swapped, carry.name, z.name)
		}
	}

	slices.Sort(swapped)
	return strings.Join(swapped, ",")
}

func main() {
	h := harness.New(solve)
	h.Run()
}
