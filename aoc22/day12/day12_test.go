package day12

import (
	"reflect"
	"testing"
)


var test_text = `
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`

var test_input = Input{
	[][]int{
		{0, 0, 1, 16, 15, 14, 13, 12},
		{0, 1, 2, 17, 24, 23, 23, 11},
		{0, 2, 2, 18, 25, 25, 23, 10},
		{0, 2, 2, 19, 20, 21, 22,  9},
		{0, 1, 3,  4,  5,  6,  7,  8},
	},
	Pos{0, 0},
	Pos{2, 5},
}

func TestParseInput(t *testing.T) {
	input := ParseInput(test_text)
	if !reflect.DeepEqual(test_input, input) {
		t.Errorf("Unexpected input:\n %v, \n expected %v", input, test_input)
	}
}

func TestPartOne(t *testing.T) {
	answer := PartOne(test_input)
	if answer != 31 {
		t.Errorf("Unexpected answer: %d", answer)
	}
}

func TestPartTwo(t *testing.T) {
	answer := PartTwo(test_input)
	if answer != 29 {
		t.Errorf("Unexpected answer: %d", answer)
	}
}
