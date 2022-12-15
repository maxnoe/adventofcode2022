package main

import (
	"github.com/maxnoe/adventofcode2022/aoc22/day01"
	"github.com/maxnoe/adventofcode2022/aoc22/day02"
	"github.com/maxnoe/adventofcode2022/aoc22/day03"
	"github.com/maxnoe/adventofcode2022/aoc22/day04"
	"github.com/maxnoe/adventofcode2022/aoc22/day05"
	"github.com/maxnoe/adventofcode2022/aoc22/day06"
	"github.com/maxnoe/adventofcode2022/aoc22/day07"
	"github.com/maxnoe/adventofcode2022/aoc22/day08"
	"github.com/maxnoe/adventofcode2022/aoc22/day09"
	"github.com/maxnoe/adventofcode2022/aoc22/day10"
	"github.com/maxnoe/adventofcode2022/aoc22/day11"
	"github.com/maxnoe/adventofcode2022/aoc22/day12"
	"github.com/maxnoe/adventofcode2022/aoc22/day13"
	"github.com/maxnoe/adventofcode2022/aoc22/day14"
	"github.com/maxnoe/adventofcode2022/aoc22/day15"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	if len(os.Args) != 2 {
		log.Fatal("Usage: ./main <day>")
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Day must be an integer: %s", err)
	}

	log.Printf("Day %d\n", day)

	switch day {
	case 1:
		day01.Day1()
	case 2:
		day02.Day2()
	case 3:
		day03.Day3()
	case 4:
		day04.Day4()
	case 5:
		day05.Day5()
	case 6:
		day06.Day6()
	case 7:
		day07.Day7()
	case 8:
		day08.Day8()
	case 9:
		day09.Day9()
	case 10:
		day10.Day10()
	case 11:
		day11.Day11()
	case 12:
		day12.Day12()
	case 13:
		day13.Day13()
	case 14:
		day14.Day14()
	case 15:
		day15.Day15()
	default:
		log.Fatalf("Unknown day: %d", day)
	}
}
