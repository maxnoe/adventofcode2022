package day10

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)

type Op int

const (
	NOOP Op = iota
	ADDX
)

type Instruction struct {
	op  Op
	arg int
}

func ParseInput(input string) []Instruction {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "noop":
			instructions[i] = Instruction{NOOP, 0}
		case "addx":
			val, err := strconv.Atoi(parts[1])
			aoc22.CheckError(err)
			instructions[i] = Instruction{ADDX, val}
		default:
			log.Fatalf("Error parsing input: %s", line)
		}
	}
	return instructions
}

func Cycles(op Op) int {
	switch op {
	case NOOP:
		return 1
	case ADDX:
		return 2
	}
	panic("not possible")
}

func Execute(instructions []Instruction) int {
	x := 1
	result := 0
	idx := 0

	// how many cycles to complete current instruction
	cycles_left := Cycles(instructions[idx].op)
	instruction := instructions[idx]

	cycle := 1
	for ; ; cycle++ {
		if (cycle-20)%40 == 0 {
			strength := cycle * x
			result += strength
		}

		pos := (cycle - 1) % 40

		if (x-1) <= pos && (pos <= x+1) {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		if pos == 39 {
			fmt.Print("\n")
		}

		cycles_left -= 1
		if cycles_left == 0 {

			if instruction.op == ADDX {
				x += instruction.arg
			}

			idx += 1
			if idx == len(instructions) {
				break
			}

			instruction = instructions[idx]
			cycles_left = Cycles(instruction.op)
		}

	}
	fmt.Print("\n")
	log.Printf("execution took %d cycles, x: %d", cycle, x)
	return result
}

func Day10() {
	input, err := aoc22.GetInput(2022, 10)
	aoc22.CheckError(err)
	log.Print("Parsing input")
	instructions := ParseInput(input)
	log.Printf("Part 1: %d", Execute(instructions))

}
