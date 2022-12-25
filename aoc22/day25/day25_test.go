package day25

import (
	"strings"
	"testing"
)


var test_input = `
1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122
`

var test_snafus = strings.Split(strings.Trim(test_input, "\n"), "\n")

var test_numbers = []int {
	1747,
	906,
	198,
	11,
	201,
	31,
	1257,
	32,
	353,
	107,
	7,
	3,
	37,
}


func TestParseInput(t *testing.T) {
	numbers := ParseInput(test_input)
	for i, expected := range test_numbers {
		if numbers[i] != expected {
			t.Errorf("Expected: %d for %s, got: %d", expected, test_snafus[i], numbers[i])
		}
	}
}

func TestToSnafu(t *testing.T) {
	for i, expected := range test_snafus {
		snafu := ToSnafu(test_numbers[i])
		if snafu != expected {
			t.Errorf("Expected: '%s' for %d, got: '%s'", expected, test_numbers[i], snafu)
		}
	}
}
