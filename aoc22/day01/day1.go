package day01

import (
	"github.com/maxnoe/adventofcode2022/aoc22"
	"log"
	"sort"
	"strconv"
	"strings"
)

func ParseInput(input string) [][]int {
	elve_strs := strings.Split(strings.Trim(input, "\n"), "\n\n")
	elves := make([][]int, len(elve_strs))

	for i, elve_str := range elve_strs {
		calories := strings.Split(elve_str, "\n")
		elves[i] = make([]int, len(calories))
		for j, value := range calories {
			value, err := strconv.Atoi(value)
			if err != nil {
				log.Panicf("Error parsing input: %s", err)
			}
			elves[i][j] = value
		}
	}

	return elves
}

func SumCalories(elves [][]int) []int {
	total_calories := make([]int, len(elves))
	for i, calories := range elves {
		total_calories[i] = aoc22.Sum(calories)
	}
	return total_calories
}

func PartOne(sorted_calories []int) int {
	return sorted_calories[len(sorted_calories)-1]
}

func PartTwo(sorted_calories []int) int {
	idx := len(sorted_calories) - 3
	return aoc22.Sum(sorted_calories[idx:])
}

func Day1() {
	input, err := aoc22.GetInput(2022, 1)
	if err != nil {
		log.Panicf("Error parsing input: %s", err)
	}
	elves := ParseInput(input)
	total_calories := SumCalories(elves)
	sort.Ints(total_calories)

	log.Printf("Part1: %d", PartOne(total_calories))
	log.Printf("Part2: %d", PartTwo(total_calories))
}
