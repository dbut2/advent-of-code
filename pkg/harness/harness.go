package harness

import (
	"embed"
	"fmt"

	"github.com/dbut2/advent-of-code/pkg/benchmark"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/strings"
	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

type PreProcessor[T any] func(string) T

func Nothing() PreProcessor[string] {
	return func(s string) string {
		return s
	}
}

func SplitSequence(seq string) PreProcessor[[]string] {
	return func(s string) []string {
		return utils.ParseInput(s, seq)
	}
}

func SplitNewlines() PreProcessor[[]string] {
	return SplitSequence("\n")
}

func SplitNewlinesWithInts() PreProcessor[[][]int] {
	return func(s string) [][]int {
		return lists.Map(SplitNewlines()(s), func(l string) []int {
			return strings.Ints(l)
		})
	}
}

type Harness[T any, U comparable] struct {
	preProcessor PreProcessor[T]
	run          func(string) U
	input        string
	Tester       test.Tester[U]
}

func New[T any, U comparable](solve func(T) U, input string, tests embed.FS, preProcessor PreProcessor[T]) *Harness[T, U] {
	run := func(s string) U {
		return solve(preProcessor(s))
	}

	h := Harness[T, U]{
		preProcessor: preProcessor,
		run:          run,
		input:        input,
	}

	h.Tester = test.Register(tests, run)

	return &h
}

func (h *Harness[T, U]) Run() {
	fmt.Println(h.run(h.input))
}

func (h *Harness[T, U]) Benchmark(cond benchmark.Condition) {
	benchmark.Run(func() {
		h.run(h.input)
	}, cond)
}
