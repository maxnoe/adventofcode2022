package day25

import (
	"log"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)


func IntPow(n, exp int) int {
	res := 1
	for i := 0; i < exp; i++ {
		res *= n;
	}
	return res
}


func ParseSnafu(snafu string) int {
	n := len(snafu)

	result := 0
	for i, c := range snafu {
		base := IntPow(5, n - i - 1)
		switch c {
		case '2': result += 2 * base
		case '1': result += base
		case '0':
		case '-': result -= base
		case '=': result -= 2 * base
		default:
			log.Panicf("Unexpected input at pos %d of input '%s'", i, snafu)
		}
	}
	return result
}

func ToBase5(value int) []int8 {
	digits := make([]int8, 0)

	n := value
	for n > 0 {
		digits = append(digits, int8(n % 5))
		n = n / 5
	}

	return digits
}

func ToSnafu(value int) string {

	one := '1'
	two := '2'
	neg_one := '-'
	neg_two := '='

	negative := value < 0
	if negative {
		value = -value
		one = '-'
		two = '='
		neg_one = '1'
		neg_two = '2'
	}

	// first convert to base 5
	base5 := ToBase5(value)
	snafu := make([]rune, 0, len(base5) + 1)

	carry := false
	for _, digit := range base5 {
		if carry {
			digit += 1
		}

		switch digit {
		case 0:
			snafu = append(snafu, '0')
			carry = false
		case 1:
			snafu = append(snafu, one)
			carry = false
		case 2:
			snafu = append(snafu, two)
			carry = false
		case 3:
			snafu = append(snafu, neg_two)
			carry = true
		case 4:
			snafu = append(snafu, neg_one)
			carry = true
		case 5:
			snafu = append(snafu, '0')
			carry = true
		}
	}

	if carry {
		snafu = append(snafu, one)
	}

	// reverse
	for i := 0; i < len(snafu) / 2; i++  {
		j := len(snafu) - i - 1
		snafu[i], snafu[j] = snafu[j], snafu[i] 
	}

	return string(snafu)
}

func ParseInput(input string) []int {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	numbers := make([]int, len(lines))
	for i, line := range lines {
		numbers[i] = ParseSnafu(line)
	}
	return numbers
}


func PartOne(numbers []int) string {
	s := 0
	for _, n := range numbers {
		s += n
	}
	return ToSnafu(s)
}

func PartTwo(numbers []int) int {
	return 0
}

func Day25() {
	log.Print("Getting Input")
	text, err := aoc22.GetInput(2022, 25)
	aoc22.CheckError(err)

	log.Print("Parsing Input")
	input := ParseInput(text)
	log.Printf("Done, %d numbers", len(input))

	log.Printf("Part 1: %s", PartOne(input))
}
