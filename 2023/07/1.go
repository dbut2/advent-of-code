package main

import (
	"embed"
	"slices"
	"sort"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/harness"
	"github.com/dbut2/advent-of-code/pkg/sti"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed *.txt
var inputs embed.FS

func main() {
	h := harness.New(solve, inputs)
	h.Expect(1, 6440)
	h.Run()
}

func solve(input string) int {
	s := utils.ParseInput(input)

	type hand struct {
		hand string
		bid  int
	}
	var hands []hand
	for _, line := range s {
		split := strings.Split(line, " ")
		hands = append(hands, hand{
			hand: split[0],
			bid:  sti.Int(split[1]),
		})
	}

	// sort hands based on score
	slices.SortFunc(hands, func(a, b hand) int {
		aScore := score(a.hand)
		bScore := score(b.hand)

		// use hand score if different
		if aScore != bScore {
			return aScore - bScore
		}
		// else find first different card and use card value as comparison
		for i := 0; i < 5; i++ {
			if a.hand[i] != b.hand[i] {
				return cards[a.hand[i]] - cards[b.hand[i]]
			}
		}
		return 0
	})

	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}
	return total
}

func score(hand string) int {
	cardCounts := map[rune]int{}
	for _, char := range hand {
		cardCounts[char]++
	}

	counts := []int{}
	for _, c := range cardCounts {
		counts = append(counts, c)
	}
	sort.Ints(counts)
	slices.Reverse(counts)

	// Five of a kind, where all five cards have the same label: AAAAA
	if slices.Equal(counts, []int{5}) {
		return 7
	}
	// Four of a kind, where four cards have the same label and one card has a different label: AA8AA
	if slices.Equal(counts, []int{4, 1}) {
		return 6
	}
	// Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
	if slices.Equal(counts, []int{3, 2}) {
		return 5
	}
	// Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
	if slices.Equal(counts, []int{3, 1, 1}) {
		return 4
	}
	// Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
	if slices.Equal(counts, []int{2, 2, 1}) {
		return 3
	}
	// One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
	if slices.Equal(counts, []int{2, 1, 1, 1}) {
		return 2
	}
	// High card, where all cards' labels are distinct: 23456
	if slices.Equal(counts, []int{1, 1, 1, 1, 1}) {
		return 1
	}
	return 0
}

// card rankings
var cards = map[uint8]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}
