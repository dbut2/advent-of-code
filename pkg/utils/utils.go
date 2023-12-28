package utils

import (
	"fmt"
	"strings"
)

func Test[T comparable](input, expected T) {
	if input != expected {
		panic(fmt.Sprintf("test failed!\nexpected: %v\ngot: %v\n", expected, input))
	}
}

func ParseInput(s string, on ...string) []string {
	sep := "\n"
	if len(on) > 0 {
		sep = on[0]
	}
	return strings.Split(strings.Trim(strings.TrimSpace(s), sep), sep)
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err.Error())
	}
	return v
}
