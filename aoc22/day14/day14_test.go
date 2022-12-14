package day14

import (
	"reflect"
	"testing"
)

var test_input = `
498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`

var test_rocks = []Rocks{
	{{498, 4}, {498, 6}, {496, 6}},
	{{503, 4}, {502, 4}, {502, 9}, {494, 9}},
}

func TestParse(t *testing.T) {
	rocks := ParseInput(test_input)
	if !reflect.DeepEqual(test_rocks, rocks) {
		t.Errorf("Error parsing input:expected:\n%v\nactual:\n%v", test_rocks, rocks)
	}
}

func TestPartOne(t *testing.T) {
	answer := PartOne(test_rocks)
	if answer != 24 {
		t.Errorf("Wrong answer: %d", answer)
	}
}

func TestPartTwo(t *testing.T) {
	answer := PartTwo(test_rocks)
	if answer != 93 {
		t.Errorf("Wrong answer: %d", answer)
	}
}
