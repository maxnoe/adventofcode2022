package day02

import (
	"reflect"
	"testing"
)

var test_input string = `A Y
B X
C Z
`

var test_inputs = []Input{
	{ROCK, PaperDraw},
	{PAPER, RockLoose},
	{SCISSORS, ScissorsWin},
}

var test_rounds = []Round{
	{ROCK, PAPER},
	{PAPER, ROCK},
	{SCISSORS, SCISSORS},
}

func TestParseInput(t *testing.T) {
	moves := ParseInput(test_input)

	if !reflect.DeepEqual(moves, test_inputs) {
		t.Errorf("Unexpected parse result")
	}
}

func TestOutcomes(t *testing.T) {
	rounds := []Round{
		{ROCK, ROCK},
		{ROCK, PAPER},
		{ROCK, SCISSORS},
		{PAPER, PAPER},
		{PAPER, SCISSORS},
		{PAPER, ROCK},
		{SCISSORS, SCISSORS},
		{SCISSORS, ROCK},
		{SCISSORS, PAPER},
	}
	outcomes := []Outcome{DRAW, WIN, LOSS, DRAW, WIN, LOSS, DRAW, WIN, LOSS}
	for i, round := range rounds {
		outcome := GetOutcome(round)
		expected := outcomes[i]
		if outcome != expected {
			t.Errorf("%d / %d should be %d, got %d", round.Opponent, round.Player, expected, outcome)
		}
	}
}

func TestPoints(t *testing.T) {
	expected := []int{8, 1, 6}
	for i, round := range test_rounds {
		if Points(round) != expected[i] {
			t.Errorf("Wrong answer for round %d", i)
		}
	}
}

func TestPartOne(t *testing.T) {
	if PartOne(test_inputs) != 15 {
		t.Error("Wrong answer")
	}
}

func TestPartTwo(t *testing.T) {
	if PartTwo(test_inputs) != 12 {
		t.Error("Wrong answer")
	}
}
