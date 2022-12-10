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

func (t Tester[T]) Expected(n int, expected T) {
	file, err := t.fs.ReadFile(fmt.Sprintf("test%d.txt", n))
	if err != nil {
		panic(err.Error())
	}
	input := string(file)
	out := t.solver(input)
	if out != expected {
		panic(fmt.Sprintf("test failed!\nexpected: %v\ngot: %v\n", expected, out))
	}
}
