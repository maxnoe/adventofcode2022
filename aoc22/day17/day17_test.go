package day17

import "testing"

var test_input = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"


func TestPartOne(t *testing.T) {
	answer := PartOne(test_input)
	if answer != 3068 {
		t.Errorf("Wrong answer: %d", answer)
	}
}
