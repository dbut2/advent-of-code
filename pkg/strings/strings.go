package strings

import (
	"github.com/dbut2/advent-of-code/pkg/buffers"
	"github.com/dbut2/advent-of-code/pkg/chars"
)

func Ints(s string) []int {
	var ints []int
	buffer := buffers.Number(0)
	set := false
	for _, char := range s {
		if chars.IsNum(char) {
			buffer.Add(chars.NumVal(char))
			set = true
			continue
		}
		if set {
			ints = append(ints, buffer.Clear())
			set = false
		}
	}
	if set {
		ints = append(ints, buffer.Clear())
	}
	return ints
}
