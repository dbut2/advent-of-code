package main

import (
	"embed"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 145)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input, ",")

	type lens struct {
		label string
		focal int
	}
	boxes := [256][]lens{}

	for _, line := range s {
		if strings.Contains(line, "-") {
			splits := strings.Split(line, "-")
			label := splits[0]

			hash := 0
			for _, char := range label {
				hash += int(char)
				hash *= 17
				hash %= 256
			}

			for i, lens := range boxes[hash] {
				if lens.label == label {
					boxes[hash] = append(boxes[hash][:i], boxes[hash][i+1:]...)
					break
				}
			}
		}

		if strings.Contains(line, "=") {
			splits := strings.Split(line, "=")
			label := splits[0]
			focal := sti.Int(splits[1])

			hash := 0
			for _, char := range label {
				hash += int(char)
				hash *= 17
				hash %= 256
			}

			l := lens{
				label: label,
				focal: focal,
			}

			replaced := false
			for i, lens := range boxes[hash] {
				if lens.label == l.label {
					boxes[hash][i] = l
					replaced = true
				}
			}
			if replaced {
				continue
			}

			boxes[hash] = append(boxes[hash], l)
		}
	}

	total := 0
	for i, box := range boxes {
		for j, lens := range box {
			total += (i + 1) * (j + 1) * lens.focal
		}
	}

	return total
}
