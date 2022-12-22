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


type Key struct {
	shape_idx int
	time_idx int
}

type Data struct {
	rock int
	time int
}


func Simulate(input string, n_rocks int) int {
	time := 0
	max_y := -1
	rocks := make(Set)

	seen_combinations := make(map[Key]Data)

	period := 0
	offset := 0
	base_height := 0
	delta_t := 0
	rock := 0
	shape_idx := rock % len(SHAPES)
	shape := SHAPES[shape_idx]

	
	for rock = 0; rock < n_rocks; rock++ {
		pos := Pos{2, max_y + 4}

		shape_idx = rock % len(SHAPES)
		shape = SHAPES[shape_idx]

		// log.Printf("Rock %d, max_y=%d, shape=%d, time=%d", rock, max_y + 1, shape_idx, time)

		key := Key{shape_idx, time % len(input)}
		data, inside := seen_combinations[key]

		for {
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
		seen_combinations[key] = Data{rock, time}


		if inside && period == 0 {
			log.Printf("Combination already seen at rock %d, %d rocks ago", data.rock, rock - data.rock)
			period = rock - data.rock
			offset = rock
			base_height = max_y
			delta_t = time - data.time
		}

		if period != 0 && rock - offset == 1000 * period {
			remaining := n_rocks - rock
			n_periods := remaining / period

			diff := max_y - base_height
			shift := diff * n_periods
			max_y += shift

			for key := range rocks {
				rocks[Pos{key.x, key.y + shift}] = true
			}
			max_y = Fill(pos, rocks, shape, max_y)

			rock += n_periods * period
			time += n_periods * delta_t
			log.Printf("Jumping ahead by %d rocks, diff=%d, total shift %d", n_periods * period, diff, shift)
			continue
		}
	}
	// log.Printf("Rock %d, max_y=%d, shape=%d, time=%d", rock, max_y + 1, shape_idx, time)
	return max_y + 1
}

func PartOne(input string) int {
	return Simulate(input, 2022)
}

func PartTwo(input string) int {
	return Simulate(input, 1_000_000_000_000)
}


func Day17() {
	log.Print("Getting Input")
	input, err := aoc22.GetInput(2022, 17)
	input = strings.Trim(strings.Trim(input,"\n"), " ")
	// log.Printf("|%s|", input)
	aoc22.CheckError(err)

	log.Print("Parsing Input")
	log.Printf("Done, %d moves", len(input))

	log.Printf("Part 1: %d", PartOne(input))
	log.Printf("Part 2: %d", PartTwo(input))
}
