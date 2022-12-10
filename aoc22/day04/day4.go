package day04

import (
	"log"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)

type Range struct {
	start int
	stop  int
}

type Instruction [2]Range

func ParseRange(s string) Range {
	parts := strings.Split(s, "-")
	val1, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Panicf("Error parsing input %s", s)
	}
	val2, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Panicf("Error parsing input %s", s)
	}
	return Range{val1, val2}
}

func ParseInput(input string) []Instruction {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")

	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		ranges := strings.Split(line, ",")
		range1 := ParseRange(ranges[0])
		range2 := ParseRange(ranges[1])
		instructions[i] = Instruction{range1, range2}
	}

	return instructions
}

func PartOne(input []Instruction) int {
	n := 0
	for _, instruction := range input {
		a := instruction[0]
		b := instruction[1]
		a_in_b := a.start >= b.start && a.stop <= b.stop
		b_in_a := b.start >= a.start && b.stop <= a.stop
		if a_in_b || b_in_a {
			n += 1
		}
	}
	return n
}

func PartTwo(input []Instruction) int {
	n := 0
	for _, instruction := range input {
		a := instruction[0]
		b := instruction[1]
		overlaps := (a.start >= b.start && a.stop <= b.stop) ||
			(b.start >= a.start && b.stop <= a.stop) ||
			(a.stop >= b.stop && a.start <= b.stop) ||
			(b.stop >= a.stop && b.start <= a.stop)
		if overlaps {
			n += 1
		}
	}
	return n
}

func Day4() {
	input_text, err := aoc22.GetInput(2022, 4)
	if err != nil {
		log.Panicf("Error parsing input: %s", err)
	}

	input := ParseInput(input_text)

	log.Printf("Part1: %d", PartOne(input))
	log.Printf("Part1: %d", PartTwo(input))
}
