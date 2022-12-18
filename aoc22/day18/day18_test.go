package day18

import (
	"reflect"
	"testing"
)

var test_input = `
2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5
`

var test_cubes = map[Pos]bool {
	{2,2,2}: true,
	{1,2,2}: true,
	{3,2,2}: true,
	{2,1,2}: true,
	{2,3,2}: true,
	{2,2,1}: true,
	{2,2,3}: true,
	{2,2,4}: true,
	{2,2,6}: true,
	{1,2,5}: true,
	{3,2,5}: true,
	{2,1,5}: true,
	{2,3,5}: true,
}

var test_cubes_small = map[Pos]bool {
	{1, 1, 1}: true,
	{2, 1, 1}: true,
}


func TestParse(t *testing.T) {
	cubes := ParseInput(test_input)
	if !reflect.DeepEqual(test_cubes, cubes) {
		t.Errorf("Parsing failed: %v", cubes)
	}
}

func TestPartOne(t *testing.T) {
	answer := PartOne(test_cubes_small)
	if answer != 10 {
		t.Errorf("Wrong answer: %d", answer)
	}

	answer = PartOne(test_cubes)
	if answer != 64 {
		t.Errorf("Wrong answer: %d", answer)
	}
}

func TestPartTwo(t *testing.T) {
	answer := PartTwo(test_cubes_small)
	if answer != 10 {
		t.Errorf("Wrong answer: %d", answer)
	}

	answer = PartTwo(test_cubes)
	if answer != 58 {
		t.Errorf("Wrong answer: %d, expected 58", answer)
	}
}
