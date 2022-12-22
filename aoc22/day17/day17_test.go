package day17

import "testing"

var test_input = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"


func TestPartOne(t *testing.T) {
	answer := PartOne(test_input)
	if answer != 3068 {
		t.Errorf("Wrong answer: %d", answer)
	}
}

func TestPartTwo(t *testing.T) {
	answer := PartTwo(test_input)
	expected := 1514285714288
	if answer != expected {
		t.Errorf("Wrong answer: %d, expected %d, delta %d", answer, expected, expected - answer)
	}
}
