package day20

import (
	"log"
	"strconv"
	"errors"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)


func ParseInput(input string) []int {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	numbers := make([]int, len(lines))
	for i, line := range lines {
		val, err := strconv.Atoi(line)
		aoc22.CheckError(err)
		numbers[i] = val
	}
	return numbers
}


func Mod(idx, n int) int {
	return (idx + n) % n
}


var NotFound = errors.New("Number not found")

func Find(number int, numbers []int) (int, error) {
	for i, n := range numbers {
		if n == number {
			return i, nil
		}
	}
	return -1, NotFound
}


func Move(number int, numbers []int) {
	pos, err := Find(number, numbers)
	aoc22.CheckError(err)
	n := number
	delta := 1
	if n < 0 {
		n = -n
		delta = -1
	}

	for i := 0; i < n; i++ {
		other := Mod(pos + delta, len(numbers))
		numbers[pos], numbers[other] = numbers[other], numbers[pos]
		pos = other
	}

}

func PartOne(input []int) int {
	numbers := make([]int, len(input))
	copy(numbers, input)

	for _, number := range input {
		Move(number, numbers)
	}

	n := len(numbers)
	zero, err := Find(0, numbers)
	aoc22.CheckError(err)
	log.Printf("0 is at %d", zero)

	answer := 0
	for i := 1000; i < 4000; i += 1000 {
		val := numbers[Mod(zero + i, n)]
		log.Printf("i = %d, val=%d", i, val)
		answer += val
	}
	return answer
}

func PartTwo(numbers []int) int {
	return 0
}

func Day20() {
	log.Print("Getting Input")
	text, err := aoc22.GetInput(2022, 20)
	aoc22.CheckError(err)

	log.Print("Parsing Input")
	input := ParseInput(text)
	log.Printf("Done, %d numbers", len(input))

	log.Printf("Part 1: %d", PartOne(input))
	log.Printf("Part 2: %d", PartTwo(input))
}
