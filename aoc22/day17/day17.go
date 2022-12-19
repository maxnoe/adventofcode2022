package day17

import (
	"fmt"
	"log"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)



const WIDTH = 7


type Shape int8
const (
	HORIZONTAL Shape = iota
	PLUS
	ANGLE
	VERTICAL
	SQUARE
)

var SHAPES = [...]Shape{HORIZONTAL, PLUS, ANGLE, VERTICAL, SQUARE}


type Pos struct {
	x int
	y int
}
type Set map[Pos]bool

var POSITIONS = map[Shape][]Pos{
	HORIZONTAL: {{0, 0}, {1, 0}, {2, 0}, {3, 0}},
	PLUS: {{1, 2}, {0, 1}, {1, 1}, {2, 1}, {1, 0}},
	ANGLE: {{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
	VERTICAL: {{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	SQUARE: {{0, 0}, {1, 0}, {0, 1}, {1, 1}},
}


func CheckCollison(pos Pos, rocks Set, shape Shape) bool {
	positions := POSITIONS[shape]
	for _, delta := range positions {
		n := Pos{pos.x + delta.x, pos.y + delta.y}
		if  n.y < 0 || n.x >= WIDTH || n.x < 0  {
			return true
		}

		_, inside := rocks[n]
		if inside {return true}
	}
	return false
}

func Fill(pos Pos, rocks Set, shape Shape, max_y int) int {
	positions := POSITIONS[shape]
	for _, delta := range positions {
		p := Pos{pos.x + delta.x, pos.y + delta.y}
		rocks[p] = true
		if p.y > max_y {
			max_y = p.y
		}
	}

	return max_y
}

func PrintRocks(rocks Set, pos Pos, shape Shape, max_y int) {
	max_y = max_y + 7
	grid := make([][]byte, max_y + 1)
	for i := range grid {
		grid[i] = make([]byte, 7)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, delta := range POSITIONS[shape] {
		grid[max_y - (pos.y + delta.y)][pos.x + delta.x] = '@'
	}
	for rock := range rocks {
		grid[max_y - rock.y][rock.x] = '#'
	}

	for _, line := range grid {
		fmt.Print("|")
		fmt.Print(string(line))
		fmt.Print("|\n")
	}
	fmt.Println("---------")

}


func PartOne(input string) int {
	n_rocks := 0
	time := 0
	max_y := -1
	rocks := make(Set)
	
	for n_rocks < 2023 {
		idx := n_rocks % len(SHAPES)
		shape := SHAPES[idx]
		pos := Pos{2, max_y + 4}

		for {
			// PrintRocks(rocks, pos, shape, max_y)
			move := input[time % len(input)]
			time++

			if move ==  '>'  {
				if !CheckCollison(Pos{pos.x + 1, pos.y}, rocks, shape) {
					pos.x++
				}
			} else {
				if !CheckCollison(Pos{pos.x - 1, pos.y}, rocks, shape) {
					pos.x--
				}
			}

			if CheckCollison(Pos{pos.x, pos.y - 1}, rocks, shape) {
				max_y = Fill(pos, rocks, shape, max_y)
				break
			}
			pos.y--
		}
		n_rocks++
	}
	return max_y
}

func PartTwo(input string) int {
	return 0
}


func Day17() {
	log.Print("Getting Input")
	input, err := aoc22.GetInput(2022, 17)
	input = strings.Trim(input,"\n")
	aoc22.CheckError(err)

	log.Print("Parsing Input")
	log.Printf("Done, %d moves", len(input))

	log.Printf("Part 1: %d", PartOne(input))
	log.Printf("Part 2: %d", PartTwo(input))
}
