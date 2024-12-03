package main

import (
	"embed"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/math"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test*.txt
var tests embed.FS

func main() {
	h := harness.New(solve, input, tests)
	h.Tester.Expect(1, 525152)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	wg := &sync.WaitGroup{}
	total := atomic.Int64{}
	for _, line := range s {
		wg.Add(1)
		springs := strings.Split(line, " ")[0]
		goals := strings.Split(line, " ")[1]

		longerSprings := strings.Repeat("?"+springs, 5)[1:]
		longerGoalNumbers := sti.Stis(strings.Split(strings.Repeat(","+goals, 5)[1:], ","))

		go func() {
			total.Add(int64(validSets(longerSprings, longerGoalNumbers)))
			wg.Done()
		}()
	}
	wg.Wait()
	return int(total.Load())
}

func validSets(springs string, goalNumbers []int) int {
	// store the count of unique counts of [2]int{count, goalI}
	// when we finish processing the full springs string we can count the amount
	// of ways to achieve count=0 and goalI=len(goalNumbers)
	// this is the number of valid sets of broken springs permutations
	lastRound := map[[2]int]int{
		[2]int{0, 0}: 1,
	}

	lastI := 0
	for {
		if lastI >= len(springs) {
			break
		}

		nextI := min(lastI+5, len(springs))

		p := problem{
			springs:     springs,
			maxI:        nextI,
			goalNumbers: goalNumbers,
			totalWant:   math.Sum(goalNumbers...),
		}

		nextRound := map[[2]int]int{}
		for state, sets := range lastRound {
			nextSets := p.validSubsets(lastI, state[0], state[1])
			for nextGoalI, nextSet := range nextSets {
				if sets*nextSet > 0 {
					nextRound[nextGoalI] += sets * nextSet
				}
			}
		}
		lastRound = nextRound
		lastI = nextI
	}

	return lastRound[[2]int{0, len(goalNumbers)}]
}

type problem struct {
	springs     string
	maxI        int
	goalNumbers []int
	totalWant   int
}

func (p *problem) validSubsets(charI int, count int, goalI int) map[[2]int]int {
	// we've processed the entire spring string
	// only valid if we've used all the goal numbers
	if charI == len(p.springs) {
		if goalI == len(p.goalNumbers) {
			return map[[2]int]int{[2]int{count, goalI}: 1}
		}
		if goalI == len(p.goalNumbers)-1 && count == p.goalNumbers[goalI] {
			return map[[2]int]int{[2]int{0, goalI + 1}: 1}
		}
		return nil
	}

	// we've processed up to the maxI
	// this is valid up until this point so return 1 count for this state
	if charI == p.maxI {
		return map[[2]int]int{[2]int{count, goalI}: 1}
	}

	switch p.springs[charI] {
	case '.':
		return p.validSubsetPeriod(charI, count, goalI)
	case '#':
		return p.validSubsetHash(charI, count, goalI)
	case '?':
		values := map[[2]int]int{[2]int{0, 0}: 0}
		a := p.validSubsetPeriod(charI, count, goalI)
		for k, v := range a {
			values[k] += v
		}
		b := p.validSubsetHash(charI, count, goalI)
		for k, v := range b {
			values[k] += v
		}
		return values
	}
	return nil
}

// validSubsetPeriod returns the total number of subsets if the current character is a period
func (p *problem) validSubsetPeriod(charI int, count int, goalI int) map[[2]int]int {
	if count != 0 {
		if p.goalNumbers[goalI] != count {
			return nil
		}
		return p.validSubsets(charI+1, 0, goalI+1)
	}

	return p.validSubsets(charI+1, 0, goalI)
}

// validSubsetHash returns the total number of subsets if the current character is a hash
func (p *problem) validSubsetHash(charI int, count int, goalI int) map[[2]int]int {
	count++

	if goalI >= len(p.goalNumbers) {
		return nil
	}

	if count > p.goalNumbers[goalI] {
		return nil
	}

	return p.validSubsets(charI+1, count, goalI)
}
