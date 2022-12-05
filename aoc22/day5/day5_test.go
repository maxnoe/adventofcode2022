package day5

import (
	"reflect"
	"testing"
)

var test_text = `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

var test_input = Input{
	[][]byte{{'Z', 'N'}, {'M', 'C', 'D'}, {'P'}},
	[]Instruction{
		{1, 2, 1}, {3, 1, 3}, {2, 2, 1}, {1, 1, 2},
	},
}


func TestParseInput(t *testing.T) {
	input := ParseInput(test_text)
	if !reflect.DeepEqual(input.stacks, test_input.stacks) {
		t.Errorf("Stacks do not match: %+v", input.stacks)
	}
	if !reflect.DeepEqual(input.instructions, test_input.instructions) {
		t.Errorf("Instructions do not match: %v", input.instructions)
	}
}
