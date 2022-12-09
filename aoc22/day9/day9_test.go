package day9

import (
	"reflect"
	"testing"
)

var test_input = `
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

var test_input2 = `
R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`

var test_moves = []Move{
	{RIGHT, 4}, {UP, 4}, {LEFT, 3}, {DOWN, 1},
	{RIGHT, 4}, {DOWN, 1}, {LEFT, 5}, {RIGHT, 2},
}


func TestParse(t *testing.T) {
	moves := ParseInput(test_input)
	if !reflect.DeepEqual(test_moves, moves) {
		t.Errorf("Parsing wrong: %v", moves)
	}
}

func TestPartOne(t *testing.T) {
	answer := PartOne(test_moves)
	if answer != 13 {
		t.Errorf("wrong answer: %v", answer)
	}
}

func TestPartTwo(t *testing.T) {
	answer := PartTwo(test_moves)
	if answer != 1 {
		t.Errorf("wrong answer: %v", answer)
	}

	test_moves2 := ParseInput(test_input2)
	answer2 := PartTwo(test_moves2)
	if answer2 != 36 {
		t.Errorf("wrong answer: %v", answer)
	}
}
