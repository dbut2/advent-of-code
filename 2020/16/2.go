package main

import (
	"embed"
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/test"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	_ = t
	fmt.Println(solve(input))
}

func solve(input string) int {
	sections := strings.Split(input, "\n\n")

	ruleNames, rules := parseRules(strings.Split(sections[0], "\n"))
	your := sti.Stis(strings.Split(strings.Split(sections[1], "\n")[1], ","))
	nearby := strings.Split(sections[2], "\n")[1:]

	var validTickets [][]int

	for _, ticket := range nearby {
		ticketFields := sti.Stis(strings.Split(ticket, ","))
		valid := true
		for _, field := range ticketFields {
			fieldValid := false
			for _, rule := range rules {
				if rule(field) {
					fieldValid = true
				}
			}
			if !fieldValid {
				valid = false
			}
		}
		if valid {
			validTickets = append(validTickets, ticketFields)
		}
	}

	validRules := lists.Fill2D(len(rules), len(rules), true)
	for _, ticket := range validTickets {
		for i, field := range ticket {
			for j, rule := range rules {
				validRules[i][j] = validRules[i][j] && rule(field)
			}
		}
	}

	validRules = simplifyTable(validRules)
	var m []int
	for _, rule := range validRules {
		index := -1
		for i, val := range rule {
			if val {
				index = i
			}
		}
		m = append(m, index)
	}

	total := 1
	for i, val := range your {
		ruleName := ruleNames[m[i]]
		if ruleName != strings.TrimPrefix(ruleName, "departure") {
			total *= val
		}
	}

	return total
}

func validTable(table [][]bool) bool {
	valid := true
	for i := range table {
		tc := 0
		for j := range table[0] {
			if table[i][j] {
				tc++
			}
		}
		if tc != 1 {
			valid = false
		}
	}
	for j := range table[0] {
		tc := 0
		for i := range table {
			if table[i][j] {
				tc++
			}
		}
		if tc != 1 {
			valid = false
		}
	}
	return valid
}

func simplifyTable(table [][]bool) [][]bool {
	for !validTable(table) {
		for i := range table {
			tc := 0
			li := -1
			for j := range table[0] {
				if table[i][j] {
					tc++
					li = j
				}
			}
			if tc == 1 {
				table = setColInBut(table, li, i)
			}
		}
		for j := range table[0] {
			tc := 0
			li := -1
			for i := range table {
				if table[i][j] {
					tc++
					li = i
				}
			}
			if tc == 1 {
				table = setRowInBut(table, li, j)
			}
		}
	}
	return table
}

func setColInBut(table [][]bool, j, i int) [][]bool {
	for x := range table {
		for y := range table[0] {
			if y == j && x != i {
				table[x][y] = false
			}
		}
	}
	return table
}

func setRowInBut(table [][]bool, i, j int) [][]bool {
	for x := range table {
		for y := range table[0] {
			if x == i && y != j {
				table[x][y] = false
			}
		}
	}
	return table
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
