package day23

import "testing"

var test_input = `
....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..
`


func TestPartOne(t *testing.T) {
	elves := ParseInput(test_input)
	answer := PartOne(elves)
	if answer != 110 {
		t.Errorf("Wrong answer: %d, expected 110", answer)
	}
}

func TestPartTwo(t *testing.T) {
	elves := ParseInput(test_input)
	answer := PartTwo(elves)
	if answer != 20 {
		t.Errorf("Wrong answer: %d, expected 20", answer)
	}
}
