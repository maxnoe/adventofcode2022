package day14

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)

type Point struct {
	x int
	y int
}

type Rocks []Point

func ParseInput(input string) []Rocks {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	rocks := make([]Rocks, len(lines))
	for i, line := range lines {
		points := strings.Split(line, " -> ")
		rocks[i] = make(Rocks, len(points))
		for j, point := range points {
			xy := strings.Split(point, ",")
			x, err := strconv.Atoi(xy[0])
			aoc22.CheckError(err)
			y, err := strconv.Atoi(xy[1])
			aoc22.CheckError(err)
			rocks[i][j] = Point{x, y}
		}
	}
	return rocks
}

func MakeGrid(rock_formations []Rocks) ([][]int8, int) {
	min_x := 1000
	max_x := 0
	max_y := 0
	for _, rocks := range rock_formations {
		for _, point := range rocks {
			if point.x < min_x {
				min_x = point.x
			}
			if point.x > max_x {
				max_x = point.x
			}
			if point.y > max_y {
				max_y = point.y
			}
		}
	}

	min_y := 0
	n_x := max_x - min_x + 1
	n_y := max_y - min_y + 2

	if (500 - min_x) < (n_y + 1) {
		min_x = 500 - (n_y + 1)
	}
	if (max_x - 500) < (n_y + 1) {
		max_x = 500 + n_y + 1
	}
	n_x = max_x - min_x + 1

	grid := make([][]int8, n_y)
	for i := range grid {
		grid[i] = make([]int8, n_x)
	}

	for _, rocks := range rock_formations {
		for i := 1; i < len(rocks); i++ {
			a := rocks[i-1]
			b := rocks[i]

			// horizontal line
			if a.y == b.y {
				if a.x > b.x {
					a, b = b, a
				}
				for x := a.x; x <= b.x; x++ {
					grid[a.y][x-min_x] = 1
				}
				// vertical line
			} else {
				if a.y > b.y {
					a, b = b, a
				}
				for y := a.y; y <= b.y; y++ {
					grid[y][a.x-min_x] = 1
				}
			}
		}
	}

	return grid, min_x
}

func PrintGrid(grid [][]int8, min_x int, pos *Point) {
	for y, line := range grid {
		fmt.Printf("% 3d", y)
		for col, cell := range line {
			if y == 0 && col+min_x == 500 {
				fmt.Print("+")
				continue
			}
			if pos != nil && col+min_x == pos.x && pos.y == y {
				fmt.Print("o")
				continue
			}
			switch cell {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("#")
			case 2:
				fmt.Print("o")
			}
		}
		fmt.Print("\n")
	}
}

func FallingSand(input []Rocks, floor bool) int {
	grid, min_x := MakeGrid(input)
	max_y := len(grid) - 1
	max_x := min_x + len(grid[0]) - 1
	n_sand := 0

outer:
	for {
		pos := Point{500, 0}
		for {

			if floor && grid[0][500-min_x] == 2 {
				break outer
			}

			if pos.y+1 > max_y {
				if floor {
					grid[pos.y][pos.x-min_x] = 2
					break
				} else {
					break outer
				}
			}

			// empty spot below, just fall
			if grid[pos.y+1][pos.x-min_x] == 0 {
				pos.y++
				continue
			}

			// falling left into the void
			if pos.x-1 < min_x {
				break outer
			}
			if grid[pos.y+1][pos.x-1-min_x] == 0 {
				pos.x--
				pos.y++
				continue
			}

			// falling right into the void
			if pos.x+1 > max_x {
				break outer
			}
			if grid[pos.y+1][pos.x+1-min_x] == 0 {
				pos.x++
				pos.y++
				continue
			}

			// rest
			grid[pos.y][pos.x-min_x] = 2
			break
		}
		n_sand++
	}
	return n_sand
}

func PartOne(input []Rocks) int {
	return FallingSand(input, false)
}

func PartTwo(input []Rocks) int {
	return FallingSand(input, true)
}

func Day14() {
	log.Print("Getting Input")
	text, err := aoc22.GetInput(2022, 14)
	aoc22.CheckError(err)

	log.Print("Parsing Input")
	input := ParseInput(text)
	log.Printf("Done, %d rock formations", len(input))

	log.Printf("Part 1: %d", PartOne(input))
	log.Printf("Part 2: %d", PartTwo(input))
}
