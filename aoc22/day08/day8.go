package day08

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)

type Trees [][]int

func ParseInput(input string) Trees {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	trees := make(Trees, len(lines))
	for i, line := range lines {
		trees[i] = make([]int, len(line))
		for j, char := range line {
			val, err := strconv.Atoi(string(char))
			aoc22.CheckError(err)
			trees[i][j] = val
		}
	}

	return trees
}

func PartOne(trees Trees) int {
	visible := make([][]bool, len(trees))
	for row, tree_row := range trees {
		visible[row] = make([]bool, len(tree_row))
	}

	n_rows := len(trees)
	n_cols := len(trees[0])

	for row := 0; row < n_rows; row++ {
		largest_height := -1
		for col := 0; col < n_cols; col++ {
			height := trees[row][col]
			if height > largest_height {
				visible[row][col] = true
				largest_height = height
			}
		}
	}

	for col := 0; col < n_cols; col++ {
		largest_height := -1
		for row := 0; row < n_rows; row++ {
			height := trees[row][col]
			if height > largest_height {
				visible[row][col] = true
				largest_height = height
			}
		}
	}

	for row := 0; row < n_rows; row++ {
		largest_height := -1
		for col := n_cols - 1; col > 0; col-- {
			height := trees[row][col]
			if height > largest_height {
				visible[row][col] = true
				largest_height = height
			}
		}
	}

	for col := 0; col < n_cols; col++ {
		largest_height := -1
		for row := n_rows - 1; row > 0; row-- {
			height := trees[row][col]
			if height > largest_height {
				visible[row][col] = true
				largest_height = height
			}
		}
	}

	n_visible := 0
	for _, row := range visible {
		for _, tree := range row {
			if tree {
				n_visible += 1
			}
		}
	}
	return n_visible
}

func ScenicScore(trees Trees, row int, col int) int {

	n_rows := len(trees)
	n_cols := len(trees[0])
	height := trees[row][col]

	// down
	n_down := 0
	for r := row + 1; r < n_rows; r++ {
		n_down += 1
		if trees[r][col] >= height {
			break
		}
	}
	// up
	n_up := 0
	for r := row - 1; r >= 0; r-- {
		n_up += 1
		if trees[r][col] >= height {
			break
		}
	}
	// right
	n_right := 0
	for c := col + 1; c < n_cols; c++ {
		n_right += 1
		if trees[row][c] >= height {
			break
		}
	}
	// left
	n_left := 0
	for c := col - 1; c >= 0; c-- {
		n_left += 1
		if trees[row][c] >= height {
			break
		}
	}
	return n_down * n_up * n_left * n_right
}

func PartTwo(trees Trees) int {
	n_rows := len(trees)
	n_cols := len(trees[0])
	best_score := 0

	for row := 0; row < n_rows; row++ {
		for col := 0; col < n_cols; col++ {
			score := ScenicScore(trees, row, col)
			if score > best_score {
				best_score = score
			}
		}
	}
	return best_score
}

func Day8() {
	input, err := aoc22.GetInput(2022, 8)
	trees := ParseInput(input)
	aoc22.CheckError(err)
	fmt.Printf("Part 1: %d\n", PartOne(trees))
	fmt.Printf("Part 2: %d\n", PartTwo(trees))
}
