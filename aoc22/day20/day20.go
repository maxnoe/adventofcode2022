package day20

import (
	"log"
	"strconv"
	"errors"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)

type Item struct {
	value int
	idx int
}


func ParseInput(input string) []Item {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	numbers := make([]Item, len(lines))
	for i, line := range lines {
		val, err := strconv.Atoi(line)
		aoc22.CheckError(err)
		numbers[i] = Item{val, i}
	}
	return numbers
}


func Mod(idx, n int) int {
	return (idx + n) % n
}


var NotFound = errors.New("Number not found")

func Find(item Item, numbers []Item) (int, error) {
	for i, n := range numbers {
		if n == item {
			return i, nil
		}
	}
	return -1, NotFound
}

func FindInt(item int, numbers []Item) (int, error) {
	for i, n := range numbers {
		if n.value == item {
			return i, nil
		}
	}
	return -1, NotFound
}


func Move(number Item, numbers []Item) {
	pos, err := Find(number, numbers)
	aoc22.CheckError(err)
	n := number.value
	delta := 1
	if n < 0 {
		n = -n
		delta = -1
	}

	max := Mod(n, len(numbers) - 1)
	for i := 0; i < max; i++ {
		other := Mod(pos + delta, len(numbers))
		numbers[pos], numbers[other] = numbers[other], numbers[pos]
		pos = other
	}

}

func GetCoords(numbers []Item) int {
	n := len(numbers)
	zero, err := FindInt(0, numbers)
	aoc22.CheckError(err)
	log.Printf("0 is at %d", zero)
	answer := 0
	for i := 1000; i < 4000; i += 1000 {
		pos := Mod(zero + i, n)
		item := numbers[pos]
		log.Printf("i = %d, pos=%d, val=%d", i, pos, item.value)
		answer += item.value
	}
	return answer
}

func PartOne(input []Item) int {
	numbers := make([]Item, len(input))
	copy(numbers, input)

	for _, number := range input {
		Move(number, numbers)
	}
	return GetCoords(numbers)
}

func PartTwo(input []Item) int {
	numbers := make([]Item, len(input))
	for i, num := range input {
		numbers[i] = Item{num.value * 811589153, num.idx}
	}

	for i := 0; i < 10; i++ {
		for _, number := range input {
			Move(Item{number.value * 811589153, number.idx}, numbers)
		}
	}
	return GetCoords(numbers)
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
