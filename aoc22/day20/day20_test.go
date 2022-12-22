package day20

import (
	"reflect"
	"testing"
)


var test_input = `
1
2
-3
3
-2
0
4
`


var test_numbers = [][]int {
	{1, 2, -3, 3, -2, 0, 4},
	{2, 1, -3, 3, -2, 0, 4},
	{1, -3, 2, 3, -2, 0, 4},
	{1, 2, 3, -2, -3, 0, 4},
	{1, 2, -2, -3, 0, 3, 4},
	{1, 2, -3, 0, 3, 4, -2},
	{1, 2, -3, 0, 3, 4, -2},
	{1, 2, -3, 4, 0, 3, -2},
}


func TestParseInput(t *testing.T) {
	numbers := ParseInput(test_input)
	if !reflect.DeepEqual(numbers, test_numbers[0]) {
		t.Errorf("Parse failed: expected:\n%v\ngot\n%v", test_numbers, numbers)
	}
}

func TestMod(t *testing.T) {
	answer := Mod(-1, 2000)
	if answer != 1999 {
		t.Errorf("Wrong answer %d", answer)
	}

	answer = Mod(1 - (-3), 7)
	if answer != 4 {
		t.Errorf("Wrong answer %d", answer)
	}
}

func TestMove(t *testing.T) {
	numbers := make([]int, len(test_numbers[0]))
	copy(numbers, test_numbers[0])

	for i, number := range numbers {
		input := make([]int, len(numbers))
		copy(input, test_numbers[i])

		Move(number, input)
		if !reflect.DeepEqual(input, test_numbers[i + 1]) {
			t.Errorf("Wrong move result, i=%d, number=%d, expected:\n%v, got:\n%v", i, number, test_numbers[i+1], input)
		}
	} 
}


func TestPartOne(t *testing.T) {
	answer := PartOne(test_numbers[0])
	if answer != 3 {
		t.Errorf("Wrong answer: %d, expected 3", answer)
	}
}
