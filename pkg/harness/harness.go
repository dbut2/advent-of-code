package harness

import (
	"embed"
	"fmt"

	"github.com/dbut2/advent-of-code/pkg/test"
)

type Harness[T comparable] struct {
	test  test.Tester[T]
	solve func(string) T
}

func New[T comparable](solve func(string) T, tests embed.FS) *Harness[T] {
	h := Harness[T]{
		solve: solve,
	}
	h.test = test.Register(tests, solve)
	return &h
}

func (h *Harness[T]) Expect(n int, value T) {
	h.test.Expect(n, value)
}

func (h *Harness[T]) Solve(input string) {
	fmt.Println(h.solve(input))
}
