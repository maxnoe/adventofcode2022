package day18

import (
	"log"
	"strconv"
	"strings"

	"github.com/edwingeng/deque/v2"
	"github.com/maxnoe/adventofcode2022/aoc22"
)

type Pos struct {
	x int
	y int
	z int
}

func ParseInput(input string) map[Pos]bool {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")

	cubes := make(map[Pos]bool, len(lines))
	for _, line := range lines {
		xyz := strings.Split(line, ",")
		x, err := strconv.Atoi(xyz[0])
		aoc22.CheckError(err)
		y, err := strconv.Atoi(xyz[1])
		aoc22.CheckError(err)
		z, err := strconv.Atoi(xyz[2])
		aoc22.CheckError(err)

		cubes[Pos{x, y, z}] = true
	}

	return cubes
}

var DIRECTIONS = []Pos {
	{-1, 0, 0}, {1, 0, 0},
	{0, -1, 0}, {0, 1, 0},
	{0, 0, -1}, {0, 0, 1},
}

func PartOne(input map[Pos]bool) int {
	exposed_sides := 0
	for cube := range input {
		exposed_sides += 6

		for _, dir := range DIRECTIONS {
			neighbor := Pos{cube.x + dir.x, cube.y + dir.y, cube.z + dir.z}
			_, exists := input[neighbor]
			if exists {
				exposed_sides--
			}
		}
	}
	return exposed_sides
}

func PartTwo(input map[Pos]bool) int {
	min := Pos{1 << 62, 1<<62, 1<<62}
	max := Pos{0, 0, 0}

	for cube := range input {
		if cube.x < min.x {min.x = cube.x}
		if cube.x > max.x {max.x = cube.x}
		if cube.y < min.y {min.y = cube.y}
		if cube.y > max.y {max.y = cube.y}
		if cube.z < min.z {min.z = cube.z}
		if cube.z > max.z {max.z = cube.z}
	}

	min.x -= 1
	min.y -= 1
	min.z -= 1
	max.x += 1
	max.y += 1
	max.z += 1

	steam := make(map[Pos]bool)
	visited := make(map[Pos]bool)
	to_check := deque.NewDeque[Pos]()
	to_check.PushBack(Pos{min.x, min.y, min.z})

	i := 0
	for to_check.Len() > 0 {
		i++
		pos := to_check.PopFront()
			
		_, been_there := visited[pos]	
		if been_there {
			continue
		}

		visited[pos] = true

		_, is_cube := input[pos]
		if is_cube {
			continue
		}

		if pos.x < min.x || pos.y < min.y || pos.z < min.z || pos.x > max.x || pos.y > max.y || pos.z > max.z {
			continue
		}


		steam[pos] = true
		for _, dir := range DIRECTIONS {
			n := Pos{pos.x + dir.x, pos.y + dir.y, pos.z + dir.z}
			_, exists := visited[n]
			if !exists {
				to_check.PushBack(n)
			}

		}
	}

	exposed_sides := 0
	for cube := range input {
		for _, dir := range DIRECTIONS {
			neighbor := Pos{cube.x + dir.x, cube.y + dir.y, cube.z + dir.z}
			_, exists := steam[neighbor]
			if exists {
				exposed_sides++
			}
		}
	}

	return exposed_sides
}

func Day18() {
	log.Print("Getting Input")
	text, err := aoc22.GetInput(2022, 18)
	aoc22.CheckError(err)

	log.Print("Parsing Input")
	input := ParseInput(text)
	log.Printf("Done, %d lava cubes", len(input))

	log.Printf("Part 1: %d", PartOne(input))
	log.Printf("Part 2: %d", PartTwo(input))
}
