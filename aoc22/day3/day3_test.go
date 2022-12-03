package day3

import (
	"testing"
)


const test_input = `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

var test_rucksacks = []Rucksack {
	{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
	{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
	{"PmmdzqPrV", "vPwwTWBwg"},
	{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"},
	{"ttgJtRGJ", "QctTZtZT"},
	{"CrZsJsPPZsGz", "wwsLwLmpwMDw"},
}

var test_duplicates = []rune("pLPvts")


func TestParseInputs(t *testing.T) {
	rucksacks := ParseInput(test_input)
	for i, expected := range test_rucksacks {
		if rucksacks[i] != expected {
			t.Errorf("Parsing did not work: %s %s", rucksacks[i].Comp1, rucksacks[i].Comp2)
		}
	}
}

func TestFindDuplicates(t *testing.T) {
	for i, rucksack := range test_rucksacks {
		duplicate := FindDuplicate(rucksack)
		if duplicate != test_duplicates[i] {
			t.Errorf("Got %s, expected %s for %d", string(duplicate), string(test_duplicates[i]), i)
		}
	}
}

func TestPartOne(t *testing.T) {
	answer := PartOne(test_rucksacks)
	if answer != 157 {
		t.Errorf("Wrong answer: %d", answer)
	}
}


func TestPartTwo(t *testing.T) {
	answer := PartTwo(test_rucksacks)
	if answer != 70 {
		t.Errorf("Wrong answer: %d", answer)
	}
}
