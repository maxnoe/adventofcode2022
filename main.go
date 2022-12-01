package main

import (
	"fmt"
	"github.com/maxnoe/adventofcode2022/aoc22/day1"
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
	}
}
