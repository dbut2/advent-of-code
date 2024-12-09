package harness

import (
	"embed"
	"fmt"
	strings2 "strings"

	"github.com/dbut2/advent-of-code/pkg/benchmark"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/space"
	"github.com/dbut2/advent-of-code/pkg/strings"
)

// Harness manages running each days' challenge code, including parsing the
// input before calling solve, and managing test expectations from examples.
type Harness[T any, U comparable] struct {
	preProcessor PreProcessor[T]
	solve        func(T) U
	inputs       embed.FS
	silent       bool
}

// HarnessOpt modifies the Harness when initialising.
type HarnessOpt[T any, U comparable] func(harness *Harness[T, U])

// WithPreProcessor explicitly sets the PreProcessor to be used.
func WithPreProcessor[T any, U comparable](preProcessor PreProcessor[T]) HarnessOpt[T, U] {
	return func(h *Harness[T, U]) {
		h.preProcessor = preProcessor
	}
}

func WithSilence[T any, U comparable]() HarnessOpt[T, U] {
	return func(h *Harness[T, U]) {
		h.silent = true
	}
}

// New returns a new Harness. At minimum a solve function and some inputs are
// required.
func New[T any, U comparable](solve func(T) U, inputs embed.FS, opts ...HarnessOpt[T, U]) *Harness[T, U] {
	h := Harness[T, U]{
		preProcessor: defaultPreProcessor[T](),
		solve:        solve,
		inputs:       inputs,
	}

	for _, opt := range opts {
		opt(&h)
	}

	return &h
}

// PreProcessor is a function that process the input data before passing to the
// solve function.
type PreProcessor[T any] func(string) T

// SplitSequence trims and splits the input on the seq string sequence
func SplitSequence(seq string) PreProcessor[[]string] {
	return func(s string) []string {
		return strings2.Split(strings2.TrimSpace(s), seq)
	}
}

// SplitNewlines is the default processing that splits the input on newlines
func SplitNewlines() PreProcessor[[]string] {
	return SplitSequence("\n")
}

// Ints process the input to a slice of ints for each line
func Ints() PreProcessor[[][]int] {
	return func(s string) [][]int {
		return lists.Map(SplitNewlines()(s), func(l string) []int {
			return strings.Ints(l)
		})
	}
}

// Grid processes the input as a grid of bytes
func Grid() PreProcessor[space.Grid[byte]] {
	return func(s string) space.Grid[byte] {
		return space.NewGridFromInput(SplitNewlines()(s))
	}
}

func DoubleSection() PreProcessor[[2]string] {
	return func(s string) [2]string {
		return [2]string(SplitSequence("\n\n")(s))
	}
}

func DoubleSectionLines() PreProcessor[[2][]string] {
	return func(s string) [2][]string {
		parts := DoubleSection()(s)
		return [2][]string{SplitNewlines()(parts[0]), SplitNewlines()(parts[1])}
	}
}

func defaultPreProcessor[T any]() PreProcessor[T] {
	switch any(*new(T)).(type) {
	case string:
		return any(PreProcessor[string](strings2.TrimSpace)).(PreProcessor[T])
	case []string:
		return any(SplitNewlines()).(PreProcessor[T])
	case [2]string:
		return any(DoubleSection()).(PreProcessor[T])
	case [2][]string:
		return any(DoubleSectionLines()).(PreProcessor[T])
	case [][]int:
		return any(Ints()).(PreProcessor[T])
	case space.Grid[byte]:
		return any(Grid()).(PreProcessor[T])
	default:
		panic("no supported preprocessor for type")
	}
}

func (h *Harness[T, U]) run(s string) U {
	return h.solve(h.preProcessor(s))
}

// Run will execute the harness with the main input.
func (h *Harness[T, U]) Run() {
	out := h.run(h.getInput())
	if !h.silent {
		fmt.Println(out)
	}
}

func (h *Harness[T, U]) readInput(s string) string {
	input, err := h.inputs.ReadFile(fmt.Sprintf("%s.txt", s))
	if err != nil {
		panic(err.Error())
	}
	return string(input)
}

func (h *Harness[T, U]) getInput() string {
	return h.readInput("input")
}

func (h *Harness[T, U]) getTestInput(n int) string {
	return h.readInput(fmt.Sprintf("test%d", n))
}

// Expect will set a required assertion on the output of solve for a test input.
func (h *Harness[T, U]) Expect(n int, expected U) {
	input := h.getTestInput(n)
	out := h.run(input)
	if out != expected {
		panic(fmt.Sprintf("test failed!\n expected: %v\ngot: %v\n", expected, out))
	}
}

// Benchmark is a utility to run a benchmark on solve using the main input.
func (h *Harness[T, U]) Benchmark(cond benchmark.Condition) {
	benchmark.Run(func() {
		h.run(h.getInput())
	}, cond)
}
