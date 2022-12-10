package day06

import (
	"fmt"
	"log"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/maxnoe/adventofcode2022/aoc22"
)

func PartOne(input string) int {
	if len(input) < 4 {
		return 0
	}

	i := 3
	for ; i < len(input); i++ {
		a := input[i-3]
		b := input[i-2]
		c := input[i-1]
		d := input[i]
		if a != b && a != c && a != d && b != c && b != d && c != d {
			return i + 1
		}
	}

	return 0
}

func PartTwo(input string) int {
	if len(input) < 14 {
		return 0
	}

	i := 14
	for ; i < len(input); i++ {
		to_check := input[i-14 : i]
		set := hashset.New()
		for _, char := range to_check {
			set.Add(char)
		}

		if set.Size() == 14 {
			return i
		}
	}

	return 0
}

func Day6() {
	input, err := aoc22.GetInput(2022, 6)
	if err != nil {
		log.Fatalf("Error getting input: %s", err)
	}
	fmt.Printf("Part1: %d\n", PartOne(input))
	fmt.Printf("Part2: %d\n", PartTwo(input))
}
