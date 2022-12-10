package day08

import (
	"reflect"
	"testing"
)

var test_input = `
30373
25512
65332
33549
35390
`

var test_trees = Trees{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

func TestParse(t *testing.T) {
	trees := ParseInput(test_input)
	if !reflect.DeepEqual(trees, test_trees) {
		t.Errorf("Wrong trees: %v, expected %v", trees, test_trees)
	}
}

func TestPartOne(t *testing.T) {
	answer := PartOne(test_trees)
	if answer != 21 {
		t.Errorf("Wrong answer: %d", answer)
	}
}

func TestPartTwo(t *testing.T) {
	score1 := ScenicScore(test_trees, 1, 2)
	if score1 != 4 {
		t.Errorf("Wrong score: %d", score1)
	}
	score2 := ScenicScore(test_trees, 3, 2)
	if score2 != 8 {
		t.Errorf("Wrong score: %d", score2)
	}

	answer := PartTwo(test_trees)
	if answer != 8 {
		t.Errorf("Wrong answer: %d", answer)
	}
}
