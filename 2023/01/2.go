package main

import (
	"embed"
	_ "embed"

	"github.com/dbut2/advent-of-code/pkg/chars"
	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(2, 281)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	total := 0
	for _, line := range s {
		firstDigit := -1
		lastDigit := 0

		for i, char := range line {
			var dig int

			found := false
			if chars.IsNum(char) {
				dig = chars.NumVal(char)
				found = true
			}
			if !found {
				dig, found = checkWords(line, i, words)
			}
			if !found {
				continue
			}

			if firstDigit == -1 {
				firstDigit = dig
			}
			lastDigit = dig
		}

		total += (firstDigit * 10) + lastDigit
	}

	return total
}

func checkWords(line string, i int, words []string) (int, bool) {
	for j, word := range words {
		if checkWord(line, i, word) {
			return j, true
		}
	}
	return 0, false
}

func checkWord(line string, i int, word string) bool {
	if i+len(word) > len(line) {
		return false
	}
	return line[i:i+len(word)] == word
}
