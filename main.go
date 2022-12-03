package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/maxnoe/adventofcode2022/aoc22/day1"
	"github.com/maxnoe/adventofcode2022/aoc22/day2"
	"github.com/maxnoe/adventofcode2022/aoc22/day3"
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
	default:
		log.Fatalf("Unknown day: %d", day)
	}
}
