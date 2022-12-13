package day13

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)


type PacketPair struct {
	first interface{}
	second interface{}
}


func ParseInput(input string) []PacketPair {
	pairs := strings.Split(strings.Trim(input, "\n"), "\n\n")
	result := make([]PacketPair, len(pairs))
	for i, pair := range pairs {
		lines := strings.Split(pair, "\n")
		result[i] = PacketPair{}
		json.Unmarshal([]byte(lines[0]), &result[i].first)
		json.Unmarshal([]byte(lines[1]), &result[i].second)
	}
	return result
}

func PartOne(pairs []PacketPair) int {
	return 0
}

func PartTwo(pairs []PacketPair) int {
	return 0
}


func Day13() {
	log.Print("Getting Input")
	text, err := aoc22.GetInput(2022, 13)
	aoc22.CheckError(err)
	

	log.Print("Parsing Input")
	input := ParseInput(text)
	log.Printf("Done")

	log.Printf("Part 1: %d", PartOne(input))
	log.Printf("Part 2: %d", PartTwo(input))
}
