package day20

import (
	"log"
	"reflect"
	"testing"

	"github.com/maxnoe/adventofcode2022/aoc22"
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
		expected := test_numbers[i + 1]

		Move(number, input)
		zi, err := Find(0, input)
		aoc22.CheckError(err)
		ze, err := Find(0, expected)
		aoc22.CheckError(err)

		for i := range numbers {
			log.Printf("%d %d", input[Mod(zi + i, len(numbers))], expected[Mod(ze + i, len(numbers))])
			if input[Mod(zi + i, len(numbers))] != expected[Mod(ze + i, len(numbers))] {
				t.Errorf("Mismatch, expected:\n%v, got:\n%v", expected, input)
			}
		}

	} 
}


func TestPartOne(t *testing.T) {
	answer := PartOne(test_numbers[0])
	if answer != 3 {
		t.Errorf("Wrong answer: %d, expected 3", answer)
	}
}
