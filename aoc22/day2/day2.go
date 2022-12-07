package day2

import (
	"log"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)

type Instruction string

const (
	RockLoose   Instruction = "X"
	PaperDraw               = "Y"
	ScissorsWin             = "Z"
)

type Move int

const (
	ROCK     Move = 1
	PAPER         = 2
	SCISSORS      = 3
)

type Input struct {
	Opponent Move
	Player   Instruction
}

type Round struct {
	Opponent Move
	Player   Move
}

type Outcome int

const (
	WIN  Outcome = 6
	DRAW         = 3
	LOSS         = 0
)

func ParseInput(input string) []Input {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	inputs := make([]Input, len(lines))

	for i, line := range lines {
		instructions := strings.Split(line, " ")

		var opponent Move
		switch instructions[0] {
		case "A":
			opponent = ROCK
		case "B":
			opponent = PAPER
		case "C":
			opponent = SCISSORS
		default:
			log.Panicf("Invalid move: %s in line %d", instructions[0], i)
		}

		player := Instruction(instructions[1])
		inputs[i] = Input{opponent, player}
	}

	return inputs
}

func GetOutcome(round Round) Outcome {
	if round.Opponent == round.Player {
		return DRAW
	}

	if round.Player == ROCK && round.Opponent == SCISSORS {
		return WIN
	}

	if round.Player == SCISSORS && round.Opponent == PAPER {
		return WIN
	}

	if round.Player == PAPER && round.Opponent == ROCK {
		return WIN
	}

	return LOSS
}

func Points(round Round) int {
	outcome := GetOutcome(round)
	return int(round.Player) + int(outcome)
}

func PartOne(inputs []Input) int {
	score := 0
	for _, input := range inputs {
		var move Move
		switch input.Player {
		case RockLoose:
			move = ROCK
		case PaperDraw:
			move = PAPER
		case ScissorsWin:
			move = SCISSORS
		}
		score += Points(Round{input.Opponent, move})
	}
	return score
}

func PartTwo(inputs []Input) int {
	score := 0
	for _, input := range inputs {
		var move Move
		switch input.Player {
		case RockLoose:
			switch input.Opponent {
			case ROCK:
				move = SCISSORS
			case PAPER:
				move = ROCK
			case SCISSORS:
				move = PAPER
			}
		case PaperDraw:
			move = input.Opponent
		case ScissorsWin:
			switch input.Opponent {
			case ROCK:
				move = PAPER
			case PAPER:
				move = SCISSORS
			case SCISSORS:
				move = ROCK
			}
		}
		score += Points(Round{input.Opponent, move})
	}
	return score
}

func Day2() {
	input, err := aoc22.GetInput(2022, 2)
	if err != nil {
		log.Panicf("Error parsing input: %s", err)
	}
	inputs := ParseInput(input)
	log.Printf("Part1: %d", PartOne(inputs))
	log.Printf("Part2: %d", PartTwo(inputs))
}
