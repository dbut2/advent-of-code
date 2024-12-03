package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Tester.Expect(1, 32000000)
	h.Run()
}

type Pulse struct {
	from, to string
	high     bool
}

type Module interface {
	Receive(pulse Pulse) []Pulse
}

type Broadcast struct {
	output []string
}

func (b *Broadcast) Receive(pulse Pulse) []Pulse {
	pulses := []Pulse{}
	for _, out := range b.output {
		pulses = append(pulses, Pulse{
			from: pulse.to,
			to:   out,
			high: pulse.high,
		})
	}
	return pulses
}

type FlipFlop struct {
	state  bool
	output []string
}

func (f *FlipFlop) Receive(pulse Pulse) []Pulse {
	if !pulse.high {
		f.state = !f.state

		pulses := []Pulse{}
		for _, out := range f.output {
			pulses = append(pulses, Pulse{
				from: pulse.to,
				to:   out,
				high: f.state,
			})
		}
		return pulses
	}
	return nil
}

type Conjunction struct {
	inputs map[string]bool
	output []string
}

func (c *Conjunction) Receive(pulse Pulse) []Pulse {
	c.inputs[pulse.from] = pulse.high

	allHigh := true
	for _, high := range c.inputs {
		if !high {
			allHigh = false
			break
		}
	}

	pulses := []Pulse{}
	for _, out := range c.output {
		pulses = append(pulses, Pulse{
			from: pulse.to,
			to:   out,
			high: !allHigh,
		})
	}
	return pulses
}

func solve(input string) int {
	s := utils.ParseInput(input)

	modules := map[string]Module{}

	for j, line := range s {
		_, _ = j, line

		line = strings.ReplaceAll(line, " ", "")
		splits := strings.Split(line, "->")

		name := splits[0]

		switch {
		case strings.Contains(name, "%"):
			name = strings.ReplaceAll(name, "%", "")
			modules[name] = &FlipFlop{
				state:  false,
				output: strings.Split(splits[1], ","),
			}
		case strings.Contains(name, "&"):
			name = strings.ReplaceAll(name, "&", "")
			modules[name] = &Conjunction{
				inputs: map[string]bool{},
				output: strings.Split(splits[1], ","),
			}
		default:
			modules[name] = &Broadcast{
				output: strings.Split(splits[1], ","),
			}
		}
	}

	for name, module := range modules {
		switch v := module.(type) {
		case *FlipFlop:
			for _, out := range v.output {
				switch w := modules[out].(type) {
				case *Conjunction:
					w.inputs[name] = false
				}
			}
		case *Conjunction:
			for _, out := range v.output {
				switch w := modules[out].(type) {
				case *Conjunction:
					w.inputs[name] = false
				}
			}
		case *Broadcast:
			for _, out := range v.output {
				switch w := modules[out].(type) {
				case *Conjunction:
					w.inputs[name] = false
				}
			}
		}

	}

	low := 0
	high := 0

	pulses := lists.NewQueue[Pulse]()

	for i := 0; i < 1000; i++ {
		pulses.Push(Pulse{
			to:   "broadcaster",
			from: "button",
			high: false,
		})

		for p := range pulses.Seq {
			if p.high {
				high++
			} else {
				low++
			}

			if _, ok := modules[p.to]; !ok {
				continue
			}
			m := modules[p.to]

			out := m.Receive(p)

			for _, o := range out {
				pulses.Push(o)
			}
		}
	}

	return low * high
}
