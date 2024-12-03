package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Tester.Expect(1, 19114)
	h.Run()
}

type Workflow struct {
	Rules    []Rule
	Accepted bool
	Next     string
}

type Rule struct {
	Category string
	Larger   bool
	Amount   int
	Accepted bool
	Next     string
}

func solve(input string) int {
	s := utils.ParseInput(input)

	workflows := map[string]Workflow{}
	parts := []map[string]int{}

	ruling := true

	for _, line := range s {
		if line == "" {
			ruling = false
			continue
		}

		if ruling {
			line = strings.ReplaceAll(line, "}", "")
			workflowSplits := strings.Split(line, "{")
			name := workflowSplits[0]
			workflowSplits = strings.Split(workflowSplits[1], ",")

			w := Workflow{}
			for _, part := range workflowSplits {
				if strings.Contains(part, ":") {
					partSplits := strings.Split(part, ":")
					r := Rule{
						Category: string(partSplits[0][0]),
						Larger:   partSplits[0][1] == '>',
						Amount:   sti.Sti(partSplits[0][2:]),
						Accepted: partSplits[1] == "A",
					}
					if len(partSplits[1]) > 1 {
						r.Next = partSplits[1]
					}
					w.Rules = append(w.Rules, r)
				} else {
					w.Accepted = part == "A"
					if len(part) > 1 {
						w.Next = part
					}
				}
			}
			workflows[name] = w
		} else {
			part := map[string]int{}

			line = strings.ReplaceAll(line, "{", "")
			line = strings.ReplaceAll(line, "}", "")

			categories := strings.Split(line, ",")
			for _, category := range categories {
				part[string(category[0])] = sti.Sti(category[2:])
			}

			parts = append(parts, part)
		}
	}

	total := 0
	for _, part := range parts {
		if accepted(workflows, "in", part) {
			for _, amount := range part {
				total += amount
			}
		}
	}
	return total
}

func accepted(workflows map[string]Workflow, workflow string, part map[string]int) bool {
	w := workflows[workflow]
	for _, rule := range w.Rules {
		partAmount := part[rule.Category]

		if rule.Larger {
			if partAmount > rule.Amount {
				if rule.Next == "" {
					return rule.Accepted
				}
				return accepted(workflows, rule.Next, part)
			}
		} else {
			if partAmount < rule.Amount {
				if rule.Next == "" {
					return rule.Accepted
				}
				return accepted(workflows, rule.Next, part)
			}
		}
	}

	if w.Next == "" {
		return w.Accepted
	}
	return accepted(workflows, w.Next, part)
}
