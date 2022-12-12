package test

import (
	"embed"
	"fmt"
)

type Tester[T comparable] struct {
	fs     embed.FS
	solver func(string) T
}

func Register[T comparable](fs embed.FS, f func(string) T) Tester[T] {
	return Tester[T]{
		fs:     fs,
		solver: f,
	}
}

func (t Tester[T]) GetTestInput(n int) string {
	input, err := t.fs.ReadFile(fmt.Sprintf("test%d.txt", n))
	if err != nil {
		panic(err.Error())
	}
	return string(input)
}

func (t Tester[T]) Expect(n int, expected T) {
	input := t.GetTestInput(n)
	out := t.solver(input)
	if out != expected {
		panic(fmt.Sprintf("test failed!\nexpected: %v\ngot: %v\n", expected, out))
	}
}
