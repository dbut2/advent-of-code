package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(1, 71)
	fmt.Println(solve(input))
}

func solve(input string) int {
	sections := strings.Split(input, "\n\n")

	_, rules := parseRules(strings.Split(sections[0], "\n"))
	_ = strings.Split(sections[1], "\n")[1]
	nearby := strings.Split(sections[2], "\n")[1:]

	invalids := 0

	for _, ticket := range nearby {
		for _, field := range sti.Stis(strings.Split(ticket, ",")) {
			valid := false
			for _, rule := range rules {
				if rule(field) {
					valid = true
				}
			}
			if !valid {
				invalids += field
			}
		}
	}

	return invalids
}

// Generate a list of names and valid functions to check a rule on a field
func parseRules(raw []string) ([]string, []func(int) bool) {
	var names []string
	var rules []func(int) bool
	for _, line := range raw {
		namerule := strings.Split(line, ": ")
		name := namerule[0]

		var ranges [][]int
		rawrules := strings.Split(namerule[1], " or ")
		ranges = append(ranges, sti.Stis(strings.Split(rawrules[0], "-")))
		ranges = append(ranges, sti.Stis(strings.Split(rawrules[1], "-")))

		rule := func(i int) bool {
			for _, r := range ranges {
				min, max := r[0], r[1]
				if i >= min && i <= max {
					return true
				}
			}
			return false
		}

		names = append(names, name)
		rules = append(rules, rule)
	}
	return names, rules
}
