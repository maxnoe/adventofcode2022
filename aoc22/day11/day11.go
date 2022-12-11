package day11

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)


type Op int8
const (
	ADD Op = iota
	MUL
)

type Operation struct {
	op Op
	arg int
	old bool
}


type Monkey struct {
	items []int
	operation Operation
	number int
	true_monkey int
	false_monkey int
}


func ParseInput(input string) []Monkey {
	monkey_texts := strings.Split(strings.Trim(input, "\n"), "\n\n")
	monkeys := make([]Monkey, len(monkey_texts))

	for monkey_idx, monkey_text := range monkey_texts {
		lines := strings.Split(strings.Trim(monkey_text, "\n"), "\n")

		// items
		levels := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
		items := make([]int, len(levels))
		for i, level := range(levels) {
			val, err := strconv.Atoi(level)
			aoc22.CheckError(err)
			items[i] = val
		}

		// operation
		parts := strings.Split(strings.Split(lines[2], ": ")[1], " ")
		var operation Operation
		if parts[3] == "*" {
			operation.op = MUL
		} else {
			operation.op = ADD
		}
		if parts[4] == "old" {
			operation.old = true
		} else {
			val, err := strconv.Atoi(parts[4])
			aoc22.CheckError(err)
			operation.arg = val
		}

		// Test
		parts = strings.Split(lines[3], " ") 
		number, err := strconv.Atoi(parts[len(parts) - 1])
		aoc22.CheckError(err)
		// true
		parts = strings.Split(lines[4], " ") 
		true_monkey, err := strconv.Atoi(parts[len(parts) - 1])
		aoc22.CheckError(err)
		// false_monkey
		parts = strings.Split(lines[5], " ") 
		false_monkey, err := strconv.Atoi(parts[len(parts) - 1])
		aoc22.CheckError(err)

		monkeys[monkey_idx] = Monkey{items, operation, number, true_monkey, false_monkey}
	}

	return monkeys
}


func ApplyOperation(op Operation, val int) int {
	if op.old {
		if op.op == ADD {
			return val + val
		}
		return val * val
	}

	if op.op == ADD {
		return val + op.arg
	} 
	return val * op.arg
}


func Turn(monkeys []Monkey, idx int, common_modulus int) int {
	inspections := 0
	monkey := &monkeys[idx]
	n := len(monkey.items)
	for i := 0; i < n; i++ {
		inspections += 1
		initial_level := monkey.items[0]
		level := initial_level

		monkey.items = monkey.items[1:]
		level = ApplyOperation(monkey.operation, level)

		if common_modulus == 1 {
			level /= 3
		} else {
			level = level % common_modulus
		}
	
		other := monkey.false_monkey
		if level % monkey.number == 0 {
			other = monkey.true_monkey
		}
		monkeys[other].items = append(monkeys[other].items, level)
	}
	return inspections
}


func MonkeyBusiness(monkeys []Monkey, rounds int) int {
	inspections := make([]int, len(monkeys))

	common_modulus := 1

	if rounds != 20 {
		for _, monkey := range monkeys {
			common_modulus *= monkey.number
		}
	}

	for round := 0; round < rounds; round++ {
		for monkey_idx := range monkeys {
			inspections[monkey_idx] += Turn(monkeys, monkey_idx, common_modulus)
		}
	}
	sort.Ints(inspections)
	n := len(inspections)
	return inspections[n - 1] * inspections[n - 2]
}

func PartOne(monkeys []Monkey) int {
	return MonkeyBusiness(monkeys, 20)
}


func PartTwo(monkeys []Monkey) int {
	return MonkeyBusiness(monkeys, 10000)
}


func Day11() {
	input, err := aoc22.GetInput(2022, 11)
	aoc22.CheckError(err)

	monkeys := ParseInput(input)
	log.Printf("Part 1: %d", PartOne(monkeys))

	monkeys = ParseInput(input)
	log.Printf("Part 2: %d", PartTwo(monkeys))
}
