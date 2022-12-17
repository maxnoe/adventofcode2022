package day16

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)

type Cave struct {
	Name string
	Flow int
	Connections []string
}

func ParseInput(input string) map[string]Cave {
	re := regexp.MustCompile(`Valve (\w{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)`)
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	caves := make(map[string]Cave)
	for _, line := range lines {
		groups := re.FindStringSubmatch(line)
		fmt.Printf("%v", groups[1:])

		name := groups[1]
		flow, err := strconv.Atoi(groups[2])
		connections := strings.Split(groups[3], ", ")
		aoc22.CheckError(err)
		caves[name] = Cave{name, flow, connections}
		fmt.Printf("%v\n", caves[name])
	}


	data, err := json.Marshal(caves)
	aoc22.CheckError(err)
	fmt.Printf("\n%v\n", string(data))
	return caves

}


func PartOne(caves map[string]Cave) int {
	return 0
}

func PartTwo(caves map[string]Cave) int {
	return 0
}

func Day16() {
	log.Print("Getting Input")
	text, err := aoc22.GetInput(2022, 16)
	aoc22.CheckError(err)

	log.Print("Parsing Input")
	input := ParseInput(text)
	log.Printf("Done, %d caves", len(input))

	log.Printf("Part 1: %d", PartOne(input))
	log.Printf("Part 2: %d", PartTwo(input))
}
