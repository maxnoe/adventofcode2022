package day22

import (
	"reflect"
	"testing"
)

var test_input = `
        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
`

var test_moves = []Move{10, R, 5, L, 5, R, 10, L, 4, R, 5, L, 5}

func TestParse(t *testing.T) {
	result := ParseInput(test_input)
	if !reflect.DeepEqual(test_moves, result.Moves) {
		t.Errorf("Moves parsed incorrectly, got:\n%v, expected:\n%v", result, test_moves)
	}
}

func TestPartOne(t *testing.T) {
	input := ParseInput(test_input)
	answer := PartOne(input)
	if answer != 6032 {
		t.Errorf("Wrong answer: %d", answer)
	}
}


func TestPartTwo(t *testing.T) {
	input := ParseInput(test_input)
	answer := PartTwo(input, 4)
	if answer != 5031 {
		t.Errorf("Wrong answer: %d", answer)
	}
}
