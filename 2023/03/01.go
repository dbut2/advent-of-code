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
	t.Expect(1, 4361)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := utils.ParseInput(input)

	var total int

	// add all numbers found
	var buffer int
	for _, line := range s {
		for _, char := range line {
			if char >= '0' && char <= '9' {
				buffer *= 10
				buffer += int(char - '0')
			} else {
				total += buffer
				buffer = 0
			}
		}

		total += buffer
		buffer = 0
	}

	// replace everything in touching cells and everything that makes up a number that touches a symbol with a period
	for i, line := range s {
		for j, char := range line {
			if (char < '0' || char > '9') && char != '.' {
				for a := i - 1; a <= i+1; a++ {
					if a < 0 || a >= len(s) {
						break
					}

					// remove all in column above in and below
					s[a] = s[a][:j] + "." + s[a][j+1:]

					// remove all the left that make a number
					b := j - 1
					for {
						if b < 0 || b >= len(s[0]) {
							break
						}

						if s[a][b] >= '0' && s[a][b] <= '9' {
							s[a] = s[a][:b] + "." + s[a][b+1:]
						} else {
							break
						}

						b--
					}

					// remove all the right that make a number
					b = j + 1
					for {
						if b < 0 || b >= len(s[0]) {
							break
						}

						if s[a][b] >= '0' && s[a][b] <= '9' {
							s[a] = s[a][:b] + "." + s[a][b+1:]
						} else {
							break
						}

						b++
					}
				}

			}
		}
	}

	// minus what's left on the board, ie everything not touching a symbol
	buffer = 0
	for _, line := range s {
		for _, char := range line {
			if char >= '0' && char <= '9' {
				buffer *= 10
				buffer += int(char - '0')
			} else {
				total -= buffer
				buffer = 0
			}
		}

		total -= buffer
		buffer = 0
	}

	return total
}
