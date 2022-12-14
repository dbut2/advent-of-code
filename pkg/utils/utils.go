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

func ParseInput(s string) []string {
	return strings.Split(strings.Trim(s, "\n"), "\n")
}
