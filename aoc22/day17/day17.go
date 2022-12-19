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
	PLUS: {     {1, 2},
		{0, 1}, {1, 1}, {2, 1},
				{1, 0},
	},
	ANGLE: {
						{2, 2},
						{2, 1},
		{0, 0}, {1, 0}, {2, 0}, 
	},
	VERTICAL: {{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	SQUARE: {
		{0, 1}, {1, 1},
		{0, 0}, {1, 0},
	},
}


func Collides(pos Pos, rocks Set, shape Shape) bool {
	for _, delta := range POSITIONS[shape] {

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
	for _, delta := range POSITIONS[shape] {

		p := Pos{pos.x + delta.x, pos.y + delta.y}
		rocks[p] = true

		if p.y > max_y {
			max_y = p.y
		}
	}

	return max_y
}

func Print(rocks Set, pos Pos, shape Shape, max_y int) {
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
	time := 0
	max_y := -1
	rocks := make(Set)
	
	for n_rocks := 0; n_rocks < 2022; n_rocks++ {
		shape := SHAPES[n_rocks % len(SHAPES)]
		pos := Pos{2, max_y + 4}

		for {
			// Print(rocks, pos, shape, max_y)
			move := input[time % len(input)]
			time++

			if move ==  '>'  {
				if !Collides(Pos{pos.x + 1, pos.y}, rocks, shape) {
					pos.x++
				}
			} else {
				if !Collides(Pos{pos.x - 1, pos.y}, rocks, shape) {
					pos.x--
				}
			}

			if Collides(Pos{pos.x, pos.y - 1}, rocks, shape) {
				max_y = Fill(pos, rocks, shape, max_y)
				break
			}
			pos.y--
		}
	}
	return max_y + 1
}

func PartTwo(input string) int {
	return 0
}


func Day17() {
	log.Print("Getting Input")
	input, err := aoc22.GetInput(2022, 17)
	input = strings.Trim(strings.Trim(input,"\n"), " ")
	log.Printf("|%s|", input)
	aoc22.CheckError(err)

	log.Print("Parsing Input")
	log.Printf("Done, %d moves", len(input))

	log.Printf("Part 1: %d", PartOne(input))
	log.Printf("Part 2: %d", PartTwo(input))
}
