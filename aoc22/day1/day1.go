package day1

import (
	"log"
	"strconv"
	"strings"
    "sort"
    "github.com/maxnoe/adventofcode2022/aoc22"
)


func ParseInput(input string) [][]int {
    elve_strs := strings.Split(strings.Trim(input, "\n"), "\n\n");
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

func Sum(values[] int) int {
    s := 0
    for _, value := range values {
        s += value
    }
    return s
}

func SumCalories(elves [][]int ) []int {
    total_calories := make([]int, len(elves))
    for i, calories := range elves {
        total_calories[i] = Sum(calories)
    }
    return total_calories
}

func PartOne (sorted_calories []int ) int {
    return sorted_calories[len(sorted_calories) - 1]
}

func PartTwo (sorted_calories []int ) int {
    idx := len(sorted_calories) - 3
    return Sum(sorted_calories[idx:])
}


func Day1() {
    input, err := aoc22.GetInput(2022, 1)
    if (err != nil) {
        log.Panicf("Error parsing input: %s", err)
    }
    elves := ParseInput(input)
    total_calories := SumCalories(elves)
    sort.Ints(total_calories)

    log.Printf("Part1: %d", PartOne(total_calories))
    log.Printf("Part1: %d", PartTwo(total_calories))
}
