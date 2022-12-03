package day3

import (
	"log"
	"strings"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/maxnoe/adventofcode2022/aoc22"
)

type Rucksack struct {
    Comp1 string
    Comp2 string
}

func ParseInput(input string) []Rucksack {
    lines := strings.Split(strings.Trim(input, "\n"), "\n")
    rucksacks := make([]Rucksack, len(lines))

    for i, line := range lines {
        n := len(line)
        rucksacks[i] = Rucksack{line[:n/2], line[n/2:]}
    }
    return rucksacks
}

func FindDuplicate(rucksack Rucksack) rune {
    comp1 := hashset.New()
    for _, char := range rucksack.Comp1 {
        comp1.Add(char)
    }

    for _, char := range rucksack.Comp2 {
        if comp1.Contains(char) {
            return char
        }
    }
    log.Panic("Did not find duplicate item")
    return 0
}

func Priority(item rune) int {
    if item >= 'a' {
        return 1 + int(item - 'a')
    }
    return 27 + int(item - 'A')
}

func PartOne(rucksacks []Rucksack) int {
    prio_sum := 0
    for _, r := range rucksacks {
        dupe := FindDuplicate(r)
        prio_sum += Priority(dupe)
    }
    return prio_sum
}

func PartTwo(rucksacks []Rucksack) int {
    return 0
}


func Day3() {
	input, err := aoc22.GetInput(2022, 3)
	if err != nil {
		log.Panicf("Error parsing input: %s", err)
	}
	inputs := ParseInput(input)
	log.Printf("Part1: %d", PartOne(inputs))
	log.Printf("Part2: %d", PartTwo(inputs))
}
