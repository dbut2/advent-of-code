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
	t := test.Register(tests, solve)
	t.Expect(1, 2713310158)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	monkeys := make(map[int]*monkey)

	inspectionCounts := make(map[int]int)

	var current *monkey

	c := 1

	for _, str := range s {
		line := strings.Split(strings.TrimLeft(str, " "), " ")

		switch line[0] {
		case "Monkey":
			name := sti.Sti(strings.Split(line[1], ":")[0])
			m := &monkey{}
			monkeys[name] = m
			inspectionCounts[name] = 0
			current = m
		case "Starting":
			numbers := sti.Stis(strings.Split(strings.Join(line[2:], ""), ","))
			current.items = numbers
		case "Operation:":
			current.operation = func(n int) int {
				number := 0
				operand := "+"

				for _, operation := range line[3:] {
					switch operation {
					case "old":
						switch operand {
						case "+":
							number += n
						case "*":
							number *= n
						}
					case "+":
						operand = operation
					case "*":
						operand = operation
					default:
						o := sti.Sti(operation)
						switch operand {
						case "+":
							number += o
						case "*":
							number *= o
						}

					}
				}

				return number
			}
		case "Test:":
			current.test = sti.Sti(line[3])
			c *= sti.Sti(line[3])
		case "If":
			switch line[1] {
			case "true:":
				current.tcase = sti.Sti(line[5])
			case "false:":
				current.fcase = sti.Sti(line[5])
			}
		}
	}

	for i := 0; i < 10000; i++ {
		inspectionCounts = doRound(monkeys, inspectionCounts, c)
	}

	var counts []int
	for _, count := range inspectionCounts {
		counts = append(counts, count)
	}

	active := math.LargestN(counts, 2)

	return active[0] * active[1]
}

func doRound(monkeys map[int]*monkey, inspections map[int]int, c int) map[int]int {
	for i := 0; i < len(monkeys); i++ {
		m := monkeys[i]
		for j := range m.items {
			inspections[i]++

			m.items[j] = m.operation(m.items[j]) % c

			if m.items[j]%m.test == 0 {
				monkeys[m.tcase].items = append(monkeys[m.tcase].items, m.items[j])
			} else {
				monkeys[m.fcase].items = append(monkeys[m.fcase].items, m.items[j])
			}
		}
		m.items = []int{}
	}
	return inspections
}

type monkey struct {
	items     []int
	operation func(int) int
	test      int
	tcase     int
	fcase     int
}
