package day5

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)


type Instruction struct {
    n int
    from int
    to int
}


type Input struct {
    stacks [][]byte 
    instructions []Instruction
}

func CheckError(err error) {
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
}

func ParseInput(input_text string) Input {
    parts := strings.Split(strings.Trim(input_text, "\n"), "\n\n")


    stack_lines := strings.Split(parts[0], "\n")
    instruction_lines := strings.Split(parts[1], "\n");

    // all but the last column are 3 wide + 1 space
    n_stacks := 1 + (len(stack_lines[0]) - 3) / 4
    input := Input{
        make([][]byte, n_stacks),
        make([]Instruction, len(instruction_lines)),
    }

    // last line of stack input is the stack numbers
    highest_stack := len(stack_lines) - 1
    for i := highest_stack - 1; i >= 0; i-- {
        for stack := 0; stack < n_stacks; stack ++ {
            col := 1 + 4 * stack;
            value := stack_lines[i][col]
            if value != ' ' {
                input.stacks[stack] = append(input.stacks[stack], value)
            }
        }
    }

    re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
    for i, line := range instruction_lines {
        groups := re.FindStringSubmatch(line)

        n, err := strconv.Atoi(groups[1])
        CheckError(err)
        from, err := strconv.Atoi(groups[2])
        CheckError(err)
        to, err := strconv.Atoi(groups[3])
        CheckError(err)

        input.instructions[i] = Instruction{n, from, to}
    }
    
    return input
}

func Move(stacks [][]byte, n int, from int, to int) [][]byte {
    from_stack := stacks[from - 1]
    to_stack := stacks[to - 1]
    for i := 0; i < n; i++ {
        top := len(from_stack) - 1
        value := from_stack[top]
        from_stack = from_stack[:top]
        to_stack = append(to_stack, value)
    }
    stacks[from - 1] = from_stack
    stacks[to - 1] = to_stack
    return stacks
}

func PartOne(input Input) string {
    for _, inst := range input.instructions {
        input.stacks = Move(input.stacks, inst.n, inst.from, inst.to)
    } 
    result := make([]byte, len(input.stacks))
    for i, stack := range input.stacks {
        top := len(stack) - 1
        result[i] = stack[top]
    }
    return string(result)
}


func Day5() {
    input_text, err := aoc22.GetInput(2022, 5)
    if err != nil {
        log.Fatalf("Error getting input: %s", err)
    }
    input := ParseInput(input_text)

    fmt.Printf("Part1: %s\n", PartOne(input))
}
