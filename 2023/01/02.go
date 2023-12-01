package main

import (
	"embed"
	_ "embed"
	"fmt"

	"github.com/dbut2/advent-of-code/pkg/test"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	t := test.Register(tests, solve)
	t.Expect(2, 281)
	fmt.Println(solve(input))
}

func solve(input string) int {
	checkWord("ancone", 3, "one")

	s := utils.ParseInput(input)

	total := 0

	for _, line := range s {
		firstDigit := 0

		firstSet := false
		lastDigit := 0

		for i, char := range line {
			var dig int

			found := false

			if char >= '0' && char <= '9' {
				dig = int(char - '0')
				found = true
			}

			if !found {
				switch {
				case checkWord(line, i, "zero"):
					found = true
					dig = 0
				case checkWord(line, i, "one"):
					found = true
					dig = 1
				case checkWord(line, i, "two"):
					found = true
					dig = 2
				case checkWord(line, i, "three"):
					found = true
					dig = 3
				case checkWord(line, i, "four"):
					found = true
					dig = 4
				case checkWord(line, i, "five"):
					found = true
					dig = 5
				case checkWord(line, i, "six"):
					found = true
					dig = 6
				case checkWord(line, i, "seven"):
					found = true
					dig = 7
				case checkWord(line, i, "eight"):
					found = true
					dig = 8
				case checkWord(line, i, "nine"):
					found = true
					dig = 9
				}
			}

			if !found {
				continue
			}

			if !firstSet {
				firstDigit = dig
				firstSet = true
			}

			lastDigit = dig
		}

		total += (firstDigit * 10) + lastDigit
	}

	return total
}

func checkWord(line string, i int, word string) bool {
	for j := range word {
		if i+j >= len(line) {
			return false
		}

		if line[i+j] != word[j] {
			return false
		}
	}

	return true
}
