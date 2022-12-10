package main

import (
	"fmt"
	"github.com/maxnoe/adventofcode2022/aoc22/day1"
	"github.com/maxnoe/adventofcode2022/aoc22/day2"
	"github.com/maxnoe/adventofcode2022/aoc22/day3"
	"github.com/maxnoe/adventofcode2022/aoc22/day4"
	"github.com/maxnoe/adventofcode2022/aoc22/day5"
	"github.com/maxnoe/adventofcode2022/aoc22/day6"
	"github.com/maxnoe/adventofcode2022/aoc22/day7"
	"github.com/maxnoe/adventofcode2022/aoc22/day8"
	"github.com/maxnoe/adventofcode2022/aoc22/day9"
	"github.com/maxnoe/adventofcode2022/aoc22/day10"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: ./main <day>")
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Day must be an integer: %s", err)
	}

	fmt.Printf("Day %d\n", day)

	switch day {
	case 1:
		day1.Day1()
	case 2:
		day2.Day2()
	case 3:
		day3.Day3()
	case 4:
		day4.Day4()
	case 5:
		day5.Day5()
	case 6:
		day6.Day6()
	case 7:
		day7.Day7()
	case 8:
		day8.Day8()
	case 9:
		day9.Day9()
	case 10:
		day10.Day10()
	default:
		log.Fatalf("Unknown day: %d", day)
	}
}
