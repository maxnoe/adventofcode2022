package day12

import (
	"log"
	"math"
	"strings"

	"github.com/edwingeng/deque/v2"
	"github.com/maxnoe/adventofcode2022/aoc22"
)


type Pos struct {
	row int
	col int
}

type Input struct {
	height [][]int
	start Pos
	end Pos
}

func ParseInput(input string) Input {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	height := make([][]int, len(lines))
	var start Pos
	var end Pos

	for row, line := range lines {
		height[row] = make([]int, len(line))

		for col, char := range line {
			if char == 'S' {
				start = Pos{row, col}
				height[row][col] = 0
			} else if char == 'E' {
				end = Pos{row, col}
				height[row][col] = int('z' - 'a')
			} else {
				height[row][col] = int(char - 'a')
			}
		}
	}
	return Input{height, start, end}
}


func ShortestPathLength(start Pos, input Input) int {
	n_rows := len(input.height)
	n_cols := len(input.height[0])

	distance := make([][]int, n_rows)
	for row := range distance {
		distance[row] = make([]int, n_cols)
		for col := range distance[row] {
			distance[row][col] = -1
		}
	}

	pos := start
	distance[pos.row][pos.col] = 0
	to_check := deque.NewDeque[Pos]()
	to_check.PushBack(pos)

	directions := []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for to_check.Len() > 0 {
		pos = to_check.PopFront()

		for _, dir := range directions {
			n := Pos{pos.row + dir.row, pos.col + dir.col}
			// outside the grid
			if n.row < 0 || n.row == n_rows || n.col < 0 || n.col == n_cols {
				continue
			}

			// already visited
			if distance[n.row][n.col] != -1 {
				continue
			}

			// too high to reach
			if (input.height[n.row][n.col] - input.height[pos.row][pos.col]) > 1 {
				continue
			}

			distance[n.row][n.col] = distance[pos.row][pos.col] + 1

			// found target position
			if n == input.end {
				return distance[n.row][n.col]
			}

			to_check.PushBack(n)
		}
	}
	return math.MaxInt
}

func PartOne(input Input) int {
	return ShortestPathLength(input.start, input)
}

func PartTwo(input Input) int {
	shortest := math.MaxInt
	for row := range input.height {
		for col := range input.height[0] {
			if input.height[row][col] != 0 {
				continue
			}
			
			path_length := ShortestPathLength(Pos{row, col}, input)
			if path_length < shortest {
				shortest = path_length
			}
		}
	}

	return shortest
}


func Day12() {
	log.Print("Getting Input")
	text, err := aoc22.GetInput(2022, 12)
	aoc22.CheckError(err)
	

	log.Print("Parsing Input")
	input := ParseInput(text)
	log.Printf("Done")

	log.Printf("Part 1: %d", PartOne(input))
	log.Printf("Part 2: %d", PartTwo(input))
}
