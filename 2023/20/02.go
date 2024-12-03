package main

import (
	"embed"
	"slices"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Run()
}

type Module interface {
	Receive(to, from int, high bool, emit func(int, []int, bool), count int)
}

type Broadcast struct {
	to []int
}

func (b *Broadcast) Receive(to, from int, high bool, emit func(int, []int, bool), count int) {
	emit(to, b.to, high)
}

type FlipFlop struct {
	state bool
	to    []int
}

func (f *FlipFlop) Receive(to, from int, high bool, emit func(int, []int, bool), count int) {
	if !high {
		f.state = !f.state
		emit(to, f.to, f.state)
	}
}

type Conjunction struct {
	states map[int]bool
	to     []int
}

func (c *Conjunction) Receive(to, from int, high bool, emit func(int, []int, bool), count int) {
	c.states[from] = high
	for _, state := range c.states {
		if !state {
			emit(to, c.to, true)
			return
		}
	}
	emit(to, c.to, false)
}

type StatefulConjunction struct {
	*Conjunction
	Receives map[int]int
}

func (s *StatefulConjunction) Receive(to, from int, high bool, emit func(int, []int, bool), count int) {
	if high {
		if _, ok := s.Receives[from]; !ok {
			s.Receives[from] = count
		}
	}

	s.states[from] = high
	for _, state := range s.states {
		if !state {
			emit(to, s.to, true)
			return
		}
	}
	emit(to, s.to, false)
}

func (s *StatefulConjunction) Continue() bool {
	if len(s.Receives) < 4 {
		return true
	}
	return false
}

func solve(input string) int {
	s := utils.ParseInput(input)

	modules := map[int]Module{}
	names := map[string]int{}
	nextName := 0

	nameNumber := func(name string) int {
		if v, ok := names[name]; ok {
			return v
		}
		number := nextName
		nextName++
		names[name] = number
		return number
	}

	for j, line := range s {
		_, _ = j, line

		line = strings.ReplaceAll(line, " ", "")
		splits := strings.Split(line, "->")

		name := splits[0]

		switch {
		case strings.Contains(name, "%"):
			name = strings.ReplaceAll(name, "%", "")
			outputs := []int{}
			for _, output := range strings.Split(splits[1], ",") {
				outputs = append(outputs, nameNumber(output))
			}
			modules[nameNumber(name)] = &FlipFlop{
				state: false,
				to:    outputs,
			}
		case strings.Contains(name, "&"):
			name = strings.ReplaceAll(name, "&", "")
			outputs := []int{}
			for _, output := range strings.Split(splits[1], ",") {
				outputs = append(outputs, nameNumber(output))
			}
			modules[nameNumber(name)] = &Conjunction{
				states: map[int]bool{},
				to:     outputs,
			}
		default:
			outputs := []int{}
			for _, output := range strings.Split(splits[1], ",") {
				outputs = append(outputs, nameNumber(output))
			}
			modules[nameNumber(name)] = &Broadcast{
				to: outputs,
			}
		}
	}

	outputs := [59][]int{}
	inputs := [59][]int{}

	for name, module := range modules {
		switch v := module.(type) {
		case *FlipFlop:
			for _, out := range v.to {
				outputs[name] = append(outputs[name], out)
				inputs[out] = append(inputs[out], name)
				switch w := modules[out].(type) {
				case *Conjunction:
					w.states[name] = false
				}
			}
		case *Conjunction:
			for _, out := range v.to {
				outputs[name] = append(outputs[name], out)
				inputs[out] = append(inputs[out], name)
				switch w := modules[out].(type) {
				case *Conjunction:
					w.states[name] = false
				}
			}
		case *Broadcast:
			for _, out := range v.to {
				outputs[name] = append(outputs[name], out)
				inputs[out] = append(inputs[out], name)
				switch w := modules[out].(type) {
				case *Conjunction:
					w.states[name] = false
				}
			}
		}

	}

	type pulse struct {
		from, to int
		high     bool
	}
	pulses := lists.NewQueue[pulse]()
	emit := func(from int, to []int, high bool) {
		for _, o := range to {
			pulses.Push(pulse{
				from: from,
				to:   o,
				high: high,
			})
		}
	}

	rx := nameNumber("rx")
	var rxParent int
	for i, module := range modules {
		switch v := module.(type) {
		case *FlipFlop:
			if slices.Contains(v.to, rx) {
				rxParent = i
			}
		case *Conjunction:
			if slices.Contains(v.to, rx) {
				rxParent = i
			}
		}
	}

	oldRxParent := modules[rxParent].(*Conjunction)
	newRxParent := &StatefulConjunction{
		Conjunction: oldRxParent,
		Receives:    map[int]int{},
	}
	modules[rxParent] = newRxParent

	count := 0
	for newRxParent.Continue() {
		count++
		pulses.Push(pulse{
			to:   nameNumber("broadcaster"),
			from: nameNumber("button"),
			high: false,
		})

		for p := range pulses.Seq {
			if p.to == rx {
				continue
			}

			m := modules[p.to]

			m.Receive(p.to, p.from, p.high, emit, count)
		}
	}

	var ns []int
	for _, n := range newRxParent.Receives {
		ns = append(ns, n)
	}
	return math.Product(ns...)
}
