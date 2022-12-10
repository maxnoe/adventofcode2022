package day09

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)

type Dir string
const (
	UP Dir = "U"
	DOWN = "D"
	LEFT = "L"
	RIGHT = "R"
)

type Move struct {
	dir Dir
	steps int
}
type Position struct {
	x int
	y int
}

func ParseInput(input string) []Move {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	moves := make([]Move, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		var dir Dir
		switch parts[0] {
		case "U": dir = UP
		case "D": dir = DOWN
		case "L": dir = LEFT
		case "R": dir = RIGHT
		default: log.Fatalf("Unexpected direction, %s", parts[0])
		}
		steps, err := strconv.Atoi(parts[1])
		aoc22.CheckError(err)
		moves[i] = Move{dir, steps}
	}
	return moves
}


func move_head(head *Position, dir Dir) {
	switch dir {
	case UP:
		head.y += 1
	case DOWN:
		head.y -= 1
	case LEFT:
		head.x -= 1
	case RIGHT:
		head.x += 1
	}
}


func move_tail(head Position, tail *Position) {
	dx, xdir := aoc22.AbsSign(head.x - tail.x)
	dy, ydir := aoc22.AbsSign(head.y - tail.y)
	total := dx + dy

	// horizontal move
	if dx > 1 && dy == 0 {
		tail.x += xdir
	} else if dy > 1 && dx == 0 {
		tail.y += ydir
	} else if total > 2 {
		tail.x += xdir
		tail.y += ydir
	}
}

func TailPositions(moves []Move, length int) int {
	visited := make(map[Position]int)
	rope := make([]Position, length)
	tail := length - 1

	fmt.Print("START\n")
	for _, move := range moves {
		for step := 0; step < move.steps; step++ {

			for i := 0; i < len(rope); i++ {
				if i == 0 {
					move_head(&rope[0], move.dir)
				} else {
					move_tail(rope[i - 1], &rope[i])
				}
				// for i, pos := range rope {
				// 	if i > 0 {fmt.Print(",")}
				// 	fmt.Printf("%d,%d", pos.x, pos.y)
				// }
				// fmt.Print("\n")
			} 

			val, _ := visited[rope[tail]]
			visited[rope[tail]] = val + 1

		}
	}
	return len(visited)
}


func PartOne(moves []Move) int {
	return TailPositions(moves, 2)
}

func PartTwo(moves []Move) int {
	return TailPositions(moves, 10)
}

func Day9() {
	input, err := aoc22.GetInput(2022, 9)
	aoc22.CheckError(err)

	moves := ParseInput(input)
	log.Printf("Part1: %d", PartOne(moves))
	log.Printf("Part2: %d", PartTwo(moves))
}

