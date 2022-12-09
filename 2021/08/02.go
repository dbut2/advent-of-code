package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	s := strings.Split(input, "\n")
	i := solve(s)
	fmt.Println(i)
}

func solve(s []string) int {
	var ssds []*ssd

	for _, line := range s {
		ss := &ssd{
			correction: map[rune]int{},
		}

		numbers := strings.Split(line, " ")

		observation := true

		for _, n := range numbers {
			if n == "|" {
				observation = false
				continue
			}

			if observation {
				ss.observations = append(ss.observations, n)
			} else {
				ss.outputs = append(ss.outputs, n)
			}
		}

		ss.correct()

		ssds = append(ssds, ss)
	}

	total := 0
	for _, ss := range ssds {
		fmt.Println(ss)
		total += ss.value()
	}
	return total
}

type ssd struct {
	observations []string
	outputs      []string

	// index points to correct value in .lights
	correction map[rune]int
}

func (s *ssd) correct() {

	var strs [10]string

	// 1, 7, 4, 8 have unique counts
	for _, t := range s.observations {
		switch len(t) {
		case 2:
			strs[1] = t
		case 3:
			strs[7] = t
		case 4:
			strs[4] = t
		case 7:
			strs[8] = t
		}
	}

	// a is in 7 not 1
	for _, char := range strs[7] {
		if !strings.ContainsRune(strs[1], char) {
			s.correction[char] = 1
		}
	}

	// 6 has len(6) and does not contain 1, missing char is c
	for _, obs := range s.observations {
		if len(obs) != 6 {
			continue
		}

		for _, char := range strs[1] {
			if !strings.ContainsRune(obs, char) {
				strs[6] = obs
				s.correction[char] = 3
			}
		}
	}

	// f is only segment of 1 in 6
	for _, char := range strs[1] {
		if strings.ContainsRune(strs[6], char) {
			s.correction[char] = 6
		}
	}

	// 3 has len(5) and contains 1
	for _, obs := range s.observations {
		if len(obs) != 5 {
			continue
		}

		contains := true
		for _, char := range strs[1] {
			if !strings.ContainsRune(obs, char) {
				contains = false
			}
		}

		if contains {
			strs[3] = obs
		}
	}

	// b is in 4 not 3
	for _, char := range strs[4] {
		if !strings.ContainsRune(strs[3], char) {
			s.correction[char] = 2
		}
	}

	// 9 is len(6) and contains 3
	for _, obs := range s.observations {
		if len(obs) != 6 {
			continue
		}

		contains := true
		for _, char := range strs[3] {
			if !strings.ContainsRune(obs, char) {
				contains = false
			}
		}

		if contains {
			strs[9] = obs
		}
	}

	// e is in 8 not 9
	for _, char := range strs[8] {
		if !strings.ContainsRune(strs[9], char) {
			s.correction[char] = 5
		}
	}

	// g in 9 not 4 not sets
	for _, char := range strs[9] {
		if !strings.ContainsRune(strs[4], char) {
			if s.correction[char] == 0 {
				s.correction[char] = 7
			}
		}
	}

	// d in 8 not sets
	for _, char := range strs[8] {
		if s.correction[char] == 0 {
			s.correction[char] = 4
		}
	}
}

func (s ssd) value() int {
	mappings := map[display]int{
		one:   1,
		two:   2,
		three: 3,
		four:  4,
		five:  5,
		six:   6,
		seven: 7,
		eight: 8,
		nine:  9,
		zero:  0,
	}

	v := 0
	for _, output := range s.outputs {
		d := display{false, false, false, false, false, false, false}

		for _, char := range output {
			d[s.correction[char]-1] = true
		}

		v *= 10
		v += mappings[d]
	}

	return v
}

type display [7]bool

var (
	one   display = display{false, false, true, false, false, true, false}
	two   display = display{true, false, true, true, true, false, true}
	three display = display{true, false, true, true, false, true, true}
	four  display = display{false, true, true, true, false, true, false}
	five  display = display{true, true, false, true, false, true, true}
	six   display = display{true, true, false, true, true, true, true}
	seven display = display{true, false, true, false, false, true, false}
	eight display = display{true, true, true, true, true, true, true}
	nine  display = display{true, true, true, true, false, true, true}
	zero  display = display{true, true, true, false, true, true, true}
)
