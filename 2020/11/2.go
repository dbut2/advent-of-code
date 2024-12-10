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
	utils.Test(solve(test), 26)
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

			for _, dir := range [][2]int{
				{-1, -1},
				{-1, 0},
				{-1, 1},
				{0, -1},
				{0, 1},
				{1, -1},
				{1, 0},
				{1, 1},
			} {
				x, y := i, j
				for {
					x += dir[0]
					y += dir[1]

					if x < 0 {
						break
					}
					if y < 0 {
						break
					}
					if x >= len(seats) {
						break
					}
					if y >= len(row) {
						break
					}

					if seats[x][y] == NoSeat {
						continue
					}

					if seats[x][y] == Occupied {
						if i == 0 && j == 3 {
						}
						occupied++
					}

					break
				}
			}

			switch seat {
			case Empty:
				if occupied == 0 {
					newSeats[i][j] = Occupied
				}
			case Occupied:
				if occupied >= 5 {
					newSeats[i][j] = Empty
				}
			}
		}
	}

	return newSeats
}
