package day13

import (
	"encoding/json"
	"log"
	"sort"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)

type PacketPair struct {
	first  []interface{}
	second []interface{}
}

func compare(lhs interface{}, rhs interface{}) int {
	switch lt := lhs.(type) {
	case float64:
		switch rt := rhs.(type) {
		case float64:
			if lt == rt {
				return 0
			} else if lt < rt {
				return -1
			} else {
				return 1
			}
		case []interface{}:
			return compare([]interface{}{lt}, rt)
		default:
			log.Panicf("Unexpected type: %v", lt)
		}
	case []interface{}:
		switch rt := rhs.(type) {
		case float64:
			return compare(lt, []interface{}{rt})
		case []interface{}:
			for i, elem := range lt {
				if i == len(rt) {
					return 1
				}

				switch compare(elem, rt[i]) {
				case 1:
					return 1
				case 0:
					continue
				case -1:
					return -1
				}
			}
			if len(lt) == len(rt) {
				return 0
			}
			return -1
		default:
			log.Panicf("Unexpected type: %v", lt)
		}
	default:
		log.Panicf("Unexpected type: %v", lt)
	}
	return 1
}

type Packets []interface{}

func (p Packets) Len() int {
	return len(p)
}

func (p Packets) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Packets) Less(i, j int) bool {
	return compare(p[i], p[j]) == -1
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

func CorrectOrder(pair PacketPair) bool {
	return compare(pair.first, pair.second) != 1
}

func PartOne(pairs []PacketPair) int {
	result := 0
	for i, pair := range pairs {
		if CorrectOrder(pair) {
			result += i + 1
		}
	}
	return result
}

func PartTwo(pairs []PacketPair) int {
	packets := make(Packets, len(pairs)*2+2)

	divider1 := []interface{}{[]interface{}{2.0}}
	divider2 := []interface{}{[]interface{}{6.0}}
	for i, pair := range pairs {
		packets[2*i] = pair.first
		packets[2*i+1] = pair.second
	}
	packets[len(packets)-2] = divider1
	packets[len(packets)-1] = divider2
	sort.Sort(packets)
	answer := 1
	for i, packet := range packets {
		if compare(packet, divider1) == 0 {
			answer *= i + 1
		} else if compare(packet, divider2) == 0 {
			answer *= i + 1
		}
	}
	return answer
}

func Day13() {
	log.Print("Getting Input")
	text, err := aoc22.GetInput(2022, 13)
	aoc22.CheckError(err)

	log.Print("Parsing Input")
	input := ParseInput(text)
	log.Printf("Done, %d pairs", len(input))

	log.Printf("Part 1: %d", PartOne(input))
	log.Printf("Part 2: %d", PartTwo(input))
}
