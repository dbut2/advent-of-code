package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/dbut2/advent-of-code/pkg/lists"
	"github.com/dbut2/advent-of-code/pkg/utils"
)

//go:embed input.txt
var input string

//go:embed test1.txt
var test string

func main() {
	utils.Test(solve(test), 37)
	fmt.Println(solve(input))
}

func solve(input string) int {
	s := strings.Split(input, "\n")

	seats := lists.Fill2D(len(s), len(s[0]), NoSeat)
	for i, row := range s {
		for j, seat := range strings.Split(row, "") {
			if seat == "L" {
				seats[i][j] = Empty
			}
		}
	}

	oldSeats := seats
	for {
		newSeats := doRound(oldSeats)
		changed := false
		occupied := 0

		for i, row := range newSeats {
			for j, seat := range row {
				if newSeats[i][j] != oldSeats[i][j] {
					changed = true
				}

				if seat == Occupied {
					occupied++
				}
			}
		}

		if !changed {
			return occupied
		}

		oldSeats = newSeats
	}
}

const (
	NoSeat int = iota
	Empty
	Occupied
)

func doRound(seats [][]int) [][]int {
	var newSeats [][]int
	for _, row := range seats {
		var newRow []int
		for _, seat := range row {
			newRow = append(newRow, seat)
		}
		newSeats = append(newSeats, newRow)
	}

	for i, row := range seats {
		for j, seat := range row {
			occupied := 0

			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x == i && y == j {
						continue
					}

					if x < 0 {
						continue
					}
					if y < 0 {
						continue
					}
					if x >= len(seats) {
						continue
					}
					if y >= len(row) {
						continue
					}

					if seats[x][y] == Occupied {
						occupied++
					}
				}
			}

			switch seat {
			case Empty:
				if occupied == 0 {
					newSeats[i][j] = Occupied
				}
			case Occupied:
				if occupied >= 4 {
					newSeats[i][j] = Empty
				}
			}
		}
	}

	return newSeats
}
