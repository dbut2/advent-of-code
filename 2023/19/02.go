package main

import (
	"embed"
	"maps"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Tester.Expect(1, 167409079868000)
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

	for _, line := range s {
		if line == "" {
			break
		}

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
	}

	total := 0

	type job struct {
		ranges map[string][2]int
		rule   string
	}
	copyJob := func(j job) job {
		return job{
			ranges: maps.Clone(j.ranges),
			rule:   j.rule,
		}
	}

	queue := lists.NewQueue[job]()
	queue.Push(job{
		ranges: map[string][2]int{
			"x": {1, 4000},
			"m": {1, 4000},
			"a": {1, 4000},
			"s": {1, 4000},
		},
		rule: "in",
	})

	for item := range queue.Seq {
		w := workflows[item.rule]

		for _, rule := range w.Rules {
			r := item.ranges[rule.Category]
			thisR := r
			nextR := r
			if rule.Larger {
				thisR[1] = rule.Amount
				nextR[0] = rule.Amount + 1
			} else {
				thisR[0] = rule.Amount
				nextR[1] = rule.Amount - 1
			}

			if thisR == nextR {
				continue
			}

			nextItem := copyJob(item)
			nextItem.ranges[rule.Category] = nextR
			item.ranges[rule.Category] = thisR

			if rule.Next != "" {
				nextItem.rule = rule.Next
				queue.Push(nextItem)
			} else {
				if rule.Accepted {
					p := 1
					for _, r := range nextItem.ranges {
						p *= (r[1] - r[0]) + 1
					}
					total += p
				}
			}
		}

		nextItem := copyJob(item)
		if w.Next != "" {
			nextItem.rule = w.Next
			queue.Push(nextItem)
		} else {
			if w.Accepted {
				p := 1
				for _, r := range nextItem.ranges {
					p *= r[1] - r[0] + 1
				}
				total += p
			}
		}
	}

	return total
}
