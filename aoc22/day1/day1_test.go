package day1

import (
	"reflect"
	"testing"
)

var input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`
var parsedInput = [][]int{
	{1000, 2000, 3000},
	{4000},
	{5000, 6000},
	{7000, 8000, 9000},
	{10000},
}

var totalCalories = []int{6000, 4000, 11000, 24000, 10000}
var totalCaloresSorted = []int{4000, 6000, 10000, 11000, 24000}

func TestParse(t *testing.T) {
	got := ParseInput(input)
	if !reflect.DeepEqual(parsedInput, got) {
		t.Errorf("Parsed input does not match expectation")
	}
}

func TestSum(t *testing.T) {
	got := SumCalories(parsedInput)

	if !reflect.DeepEqual(totalCalories, got) {
		t.Errorf("totalCalories do not match expectation")
	}
}

func TestPart1(t *testing.T) {
	got := PartOne(totalCaloresSorted)

	if got != 24000 {
		t.Errorf("Wrong answer in part1: %d", got)
	}
}

func TestPart2(t *testing.T) {
	got := PartTwo(totalCaloresSorted)

	if got != 45000 {
		t.Errorf("Wrong answer in part2: %d", got)
	}
}
