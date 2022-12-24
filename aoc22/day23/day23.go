package day23

import (
	"fmt"
	"log"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)

type Pos struct {
	row int
	col int
}


type Set map[Pos]bool


func ParseInput(input string) Set {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	elves := make(Set)

	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				elves[Pos{row, col}] = true
			}
		}
	}
	return elves

}

type Direction int8
const (
	NORTH Direction = iota
	SOUTH
	WEST
	EAST
)

var CHECK_DIRECTIONS = [4][3]Pos {
	{{-1, -1}, {-1, 0}, {-1, 1}},
	{{1, -1}, {1, 0}, {1, 1}},
	{{-1, -1}, {0, -1}, {1, -1}},
	{{-1, 1}, {0, 1}, {1, 1}},
}

var DIRECTIONS = [4]Pos {
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}


func IsFree(pos Pos, elves Set, direction Direction) bool {
	for _, d := range CHECK_DIRECTIONS[direction] {

		_, inside := elves[Pos{pos.row + d.row, pos.col + d.col}]
		if inside {
			return false
		}
	}
	return true
}


func Print(s Set) {
	min_row := 1 << 62
	min_col := 1 << 62
	max_row := 0
	max_col := 0
	for pos := range s {
		if pos.row < min_row {min_row = pos.row}
		if pos.row > max_row {max_row = pos.row}
		if pos.col < min_col {min_col = pos.col}
		if pos.col > max_col {max_col = pos.col}
	}

	fmt.Print("\n")
	for row := min_row; row <= max_row; row ++ {
		for col := min_col; col <= max_col; col ++ {
			_, inside := s[Pos{row, col}]
			if inside {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func Simulate(input Set, max_moves int) int {
	elves := make(Set)
	for k, v := range input {
		elves[k] = v
	}

	elves_moved := len(elves)
	n_elves := len(elves)

	first_direction := NORTH
	Print(elves)

	round := 0

	for round = 1; round <= max_moves; round++ {
		proposed_moves := make(map[Pos][]Pos)

		outer:
		for elve := range elves {
			var free [4]bool
			all_free := true
			for i := 0; i < 4; i++ {
				direction := Direction((int(first_direction) + i) % 4)
				free[i] = IsFree(elve, elves, direction)
				all_free = all_free && free[i]
			}

			if all_free {
				continue
			}

			for i := 0; i < 4; i++ {
				if !free[i] {
					continue
				}

				direction := Direction((int(first_direction) + i) % 4)
				d := DIRECTIONS[direction]
				pos := Pos{elve.row + d.row, elve.col + d.col}
				proposed_moves[pos] = append(proposed_moves[pos], elve)
				continue outer
			}
		}

		first_direction = Direction((int(first_direction) + 1) % 4)

		for _, old_pos := range proposed_moves {
			for _, e := range old_pos {
				delete(elves, e)
			}
		}

		elves_moved = 0
		for new_pos, old_pos := range proposed_moves {
			if len(old_pos) > 1 {
				for _, e := range old_pos {
					elves[e] = true
				}
			} else {
				elves[new_pos] = true
				elves_moved++
			}

		}

		if len(elves) != n_elves {
			log.Panicf("number of elves changed: %d to %d", n_elves, len(elves))
		}

		log.Printf("Round %d: elves_moved: %d", round, elves_moved)

		if elves_moved == 0 {
			Print(elves)
			return round
		}
	}

	min_row := 1 << 62
	min_col := 1 << 62
	max_row := 0
	max_col := 0
	for pos := range elves {
		if pos.row < min_row {min_row = pos.row}
		if pos.row > max_row {max_row = pos.row}
		if pos.col < min_col {min_col = pos.col}
		if pos.col > max_col {max_col = pos.col}
	}

	return (max_col - min_col + 1) * (max_row - min_row + 1) - len(elves)
}

func PartOne(elves Set) int {
	return Simulate(elves, 10)
}

func PartTwo(elves Set) int {
	return Simulate(elves, 100000)
}

func Day23() {
	log.Print("Getting Input")
	text, err := aoc22.GetInput(2022, 23)
	aoc22.CheckError(err)

	log.Print("Parsing Input")
	input := ParseInput(text)
	log.Printf("Done, %d elves", len(input))

	log.Printf("Part 1: %d", PartOne(input))
	log.Printf("Part 2: %d", PartTwo(input))
}
