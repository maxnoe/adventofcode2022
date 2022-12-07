package day6

import "testing"

var test_inputs = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	"bvwbjplbgvbhsrlpgdmjqwftvncz",
	"nppdvjthqldpwncqszvftbrmjlhg:",
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}

var expectedOne = []int{7, 5, 6, 10, 11}
var expectedTwo = []int{19, 23, 23, 29, 26}

func TestPartOne(t *testing.T) {
	for i, input := range test_inputs {
		result := PartOne(input)
		if result != expectedOne[i] {
			t.Errorf("Expected %d, got %d for %s", expectedOne[i], result, input)
		}
	}
}

func TestPartTwo(t *testing.T) {
	for i, input := range test_inputs {
		result := PartTwo(input)
		if result != expectedTwo[i] {
			t.Errorf("Expected %d, got %d for %s", expectedOne[i], result, input)
		}
	}
}
