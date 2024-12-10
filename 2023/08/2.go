package main

import (
	"embed"
	"slices"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(2, 6)
	h.Run()
}

type node struct {
	value       string
	left, right string

	// z points between next full instruction cycle
	zPoints []int
	// z points up to and including first iteration of the loop
	bigZPoints []int

	// loop definition as a count of full instruction
	loopStart, loopSize int
	// loop definition as a count of individual moves
	start, size int

	next *node
}

func solve(input string) int {
	s := utils.ParseInput(input)

	var instruction string
	nodes := map[string]*node{}
	var startingNodes []*node

	for i, line := range s {
		if i == 0 {
			instruction = line
			continue
		}
		if line == "" {
			continue
		}

		line = strings.ReplaceAll(line, "=", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ")", "")

		parts := strings.Split(line, " ")
		n := &node{
			value: parts[0],
			left:  parts[2],
			right: parts[3],
		}
		nodes[n.value] = n
		if n.value[2] == 'A' {
			startingNodes = append(startingNodes, n)
		}
	}

	// calculate next node after full instruction has ran
	for _, node := range nodes {
		current := node
		for _, char := range instruction {
			switch char {
			case 'L':
				current = nodes[current.left]
			case 'R':
				current = nodes[current.right]
			}
		}
		node.next = current
	}

	// calculate loop definitions
	for _, node := range nodes {
		firstSeen := map[string]int{}
		var duplicateFirst int
		var duplicateLast int

		current := node
		for i := 0; ; i++ {
			if firstSeenAt, ok := firstSeen[current.value]; ok {
				duplicateFirst = firstSeenAt
				duplicateLast = i

				break
			}

			firstSeen[current.value] = i
			current = current.next
		}

		node.loopStart = duplicateFirst
		node.loopSize = duplicateLast - duplicateFirst

		node.start = node.loopStart * len(instruction)
		node.size = node.loopSize * len(instruction)
	}

	// calculate z points until next instruction
	for _, node := range nodes {
		current := node
		for j, char := range instruction {
			switch char {
			case 'L':
				current = nodes[current.left]
			case 'R':
				current = nodes[current.right]
			}
			if current.value[2] == 'Z' {
				node.zPoints = append(node.zPoints, j+1)
			}
		}
	}

	// calculate z points until first loop iteration
	for _, node := range nodes {
		current := node
		for i := 0; i < node.loopStart+node.loopSize; i++ {
			for _, newZPoint := range current.zPoints {
				node.bigZPoints = append(node.bigZPoints, newZPoint+(i*len(instruction)))
			}
			current = current.next
		}
	}

	loopStart := 0
	loopSize := 1
	var zPoints []int
	for _, node := range startingNodes {
		// define new loop range
		// all nodes will loop inside this
		newLoopStart := max(loopStart, node.start)
		newLoopSize := math.LCM(loopSize, node.size)

		// only use nodes z points that are inside loop
		newZPoints := lists.Filter(node.bigZPoints, func(i int) bool {
			return i >= node.start
		})
		// change z points to be their distance from loop start
		newZPoints = lists.Map(newZPoints, func(z int) int {
			return (z - newLoopStart + node.size) % node.size
		})

		if zPoints == nil {
			zPoints = newZPoints
			loopStart = newLoopStart
			loopSize = newLoopSize
		}

		// change z points to be their distance from new loop start
		zPoints = lists.Map(zPoints, func(z int) int {
			return (z + loopStart - newLoopStart + loopSize) % loopSize
		})
		// expand z points to cover new loop range
		zPoints = expand(zPoints, loopSize, newLoopSize)

		// only keep existing z points if they will collide with nodes z points
		zPoints = lists.Filter(zPoints, func(i int) bool {
			return slices.Contains(newZPoints, i%node.size)
		})

		loopStart = newLoopStart
		loopSize = newLoopSize
	}

	return slices.Min(zPoints) + loopStart
}

func expand(s []int, oldSize int, newSize int) []int {
	m := make([]int, 0, (newSize/oldSize)*len(s))
	for _, v := range s {
		for x := v; x < newSize; x += oldSize {
			m = append(m, x)
		}
	}
	return m
}
