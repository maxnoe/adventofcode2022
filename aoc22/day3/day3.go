package day3

import (
	"log"
	"strings"

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/maxnoe/adventofcode2022/aoc22"
)

func ParseInput(input string) []string {
	return strings.Split(strings.Trim(input, "\n"), "\n")
}

func FindDuplicate(rucksack string) rune {
	comp1 := hashset.New()
	n := len(rucksack) / 2
	for _, char := range rucksack[:n] {
		comp1.Add(char)
	}

	for _, char := range rucksack[n:] {
		if comp1.Contains(char) {
			return char
		}
	}
	log.Panic("Did not find duplicate item")
	return 0
}

func Priority(item rune) int {
	if item >= 'a' {
		return 1 + int(item-'a')
	}
	return 27 + int(item-'A')
}

func PartOne(rucksacks []string) int {
	prio_sum := 0
	for _, r := range rucksacks {
		dupe := FindDuplicate(r)
		prio_sum += Priority(dupe)
	}
	return prio_sum
}

func toSet(r string) *hashset.Set {
	items := hashset.New()
	for _, item := range r {
		items.Add(item)
	}
	return items
}

func PartTwo(rucksacks []string) int {
	n_elves := len(rucksacks)
	n_groups := n_elves / 3

	prio_sum := 0

	for group := 0; group < n_groups; group++ {
		first := 3 * group
		items1 := toSet(rucksacks[first])
		items2 := toSet(rucksacks[first+1])
		items3 := toSet(rucksacks[first+2])

		intersection := items1.Intersection(items2)
		intersection = intersection.Intersection(items3)

		badge := intersection.Values()[0].(rune)
		prio_sum += Priority(badge)
	}
	return prio_sum
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
