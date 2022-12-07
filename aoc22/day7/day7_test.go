package day7

import "testing"

var test_input = `
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`

func TestParseInput(t *testing.T) {
	inodes := ParseInput(test_input)
	if len(inodes) != 14 {
		t.Error("wrong number of inodes")
	}
}

func TestPartOne(t *testing.T) {
	inodes := ParseInput(test_input)
	part1 := PartOne(inodes)
	if part1 != 95437 {
		t.Errorf("wrong answer: %d", part1)
	}
}

func TestPartTwo(t *testing.T) {
	inodes := ParseInput(test_input)
	part2 := PartTwo(inodes)
	if part2 != 24933642 {
		t.Errorf("wrong answer: %d", part2)
	}
}
