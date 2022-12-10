package utils

import (
	"embed"
	"fmt"
	"strings"
)

func Test[T comparable](input, expected T) {
	if input != expected {
		panic(fmt.Sprintf("test failed!\nexpected: %v\ngot: %v\n", expected, input))
	}
}

func Test2[T comparable](fs embed.FS, f func(string) T, n int, expected T) {
	file, err := fs.ReadFile(fmt.Sprintf("test%d.txt", n))
	if err != nil {
		panic(err.Error())
	}
	input := string(file)
	out := f(input)
	if out != expected {
		panic(fmt.Sprintf("test failed!\nexpected: %v\ngot: %v\n", expected, out))
	}
}

func ParseInput(s string) []string {
	return strings.Split(s, "\n")
}
