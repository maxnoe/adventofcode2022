package day07

import (
	"log"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)

type EntryType int

const (
	File EntryType = iota
	Dir
)

type Entry struct {
	name     string
	etype    EntryType
	size     int
	children map[string]int
	parent   int
}

func ParseInput(input string) []Entry {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")

	inodes := make([]Entry, 1)
	inodes[0] = Entry{"/", Dir, -1, make(map[string]int), -1}
	cwd := 0

	if lines[0] != "$ cd /" {
		log.Fatalf("First command not starting at /, %s", lines[0])
	}

	i := 1
	level := 0
	for i < len(lines) {
		line := lines[i]

		if line == "$ ls" {
			i += 1
			line = lines[i]
			for line[0] != '$' {

				parts := strings.Split(line, " ")
				if parts[0] == "dir" {
					inodes[cwd].children[parts[1]] = len(inodes)
					inodes = append(inodes, Entry{parts[1], Dir, -1, make(map[string]int), cwd})
				} else {
					size, err := strconv.Atoi(parts[0])
					aoc22.CheckError(err)
					inodes[cwd].children[parts[1]] = len(inodes)
					inodes = append(inodes, Entry{parts[1], File, size, nil, cwd})
				}
				i += 1
				if i >= len(lines) {
					break
				}
				line = lines[i]
			}
		} else if line == "$ cd .." {
			level -= 1
			cwd = inodes[cwd].parent
			i += 1
		} else if line[:4] == "$ cd" {
			level += 1
			cwd = inodes[cwd].children[line[5:]]
			i += 1
		} else {
			log.Fatalf("Unexpected line '%s'", line)
		}
	}

	return inodes
}

func CalcSize(index int, inodes []Entry) int {
	entry := inodes[index]
	if entry.size != -1 {
		return entry.size
	}

	size := 0
	for _, child_index := range entry.children {
		size += CalcSize(child_index, inodes)
	}
	entry.size = size
	return size

}

func PartOne(inodes []Entry) int {
	total_size := 0
	for i, entry := range inodes {
		if entry.etype == Dir {
			size := CalcSize(i, inodes)
			if size < 100000 {
				total_size += size
			}
		}
	}
	return total_size
}

func PartTwo(inodes []Entry) int {
	required_space := 30000000
	disk_size := 70000000
	smallest_size := disk_size
	in_use := CalcSize(0, inodes)
	to_free := required_space - (disk_size - in_use)

	for i, entry := range inodes {
		if entry.etype == Dir {
			size := CalcSize(i, inodes)
			if size > to_free && size < smallest_size {
				smallest_size = size
			}
		}
	}
	return smallest_size
}

func Day7() {
	input, err := aoc22.GetInput(2022, 7)
	if err != nil {
		log.Panicf("Error parsing input: %s", err)
	}
	inputs := ParseInput(input)
	log.Printf("Part1: %d", PartOne(inputs))
	log.Printf("Part2: %d", PartTwo(inputs))
}
