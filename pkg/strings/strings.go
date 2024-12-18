package strings

import (
	"github.com/dbut2/advent-of-code/pkg/buffers"
	"github.com/dbut2/advent-of-code/pkg/chars"
)

// Ints will search for all sequences of digits optionally prepended by a
// negative sign (`-`) and return a list of ints.
// String to int conversion happens through buffered ascii value conversions,
// this is probably faster and safer than builtins.
func Ints(s string) []int {
	var ints []int
	buffer := buffers.Number(0)
	set := false
	neg := false
	for _, char := range s {
		if chars.IsNum(char) {
			buffer.Add(chars.NumVal(char))
			set = true
			continue
		}
		if char == '-' {
			if set {
				b := buffer.Clear()
				if neg {
					b = -b
				}
				ints = append(ints, b)
				set = false
			}
			neg = true
			continue
		}
		if set {
			b := buffer.Clear()
			if neg {
				b = -b
			}
			neg = false
			ints = append(ints, b)
			set = false
		}
	}
	if set {
		b := buffer.Clear()
		if neg {
			b = -b
		}
		ints = append(ints, b)
	}
	return ints
}
