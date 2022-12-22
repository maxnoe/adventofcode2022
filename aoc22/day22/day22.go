package day22

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)

type Tile int8

const (
	VOID Tile = 0
	MOVABLE = 1
	WALL = 2
)

type Facing int8 
const (
	RIGHT Facing = 0
	DOWN = 1
	LEFT = 2
	UP = 3
)


type Turn byte
const (
	R Turn = 'R'
	L Turn = 'L'
)

func (t Turn) String() string {
	return string(t)
}

func (f Facing) String() string {
	switch f {
	case RIGHT:
		return "RIGHT"
	case LEFT: 
		return "LEFT"
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	default:
		return "UNKNOWN"
	}
}


type Move interface{}

type Input struct {
	Board [][]Tile
	Moves []Move
}

type Pos struct {
	row int
	col int
}

type State struct {
	Pos Pos
	Facing Facing
}

var DIRECTIONS = [...]Pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func ParseInput(text string) Input {
	parts := strings.Split(strings.Trim(text, "\n"), "\n\n")

	lines := strings.Split(parts[0], "\n")

	input := Input{make([][]Tile, len(lines)), make([]Move, 0)}

	n_cols := 0
	for _, line := range lines {
		if len(line) > n_cols {n_cols = len(line)}
	}

	for i, line := range lines {
		input.Board[i] = make([]Tile, n_cols)
		for j, c := range line {
			switch c {
			case ' ': input.Board[i][j] = VOID
			case '.': input.Board[i][j] = MOVABLE
			case '#': input.Board[i][j] = WALL
			}
			
		}
	}

	re := regexp.MustCompile(`(\d+)|[LR]`)
	for _, m := range re.FindAllString(parts[1], -1) {
		if m == "L" {
			input.Moves = append(input.Moves, L)
		} else if m == "R" {
			input.Moves = append(input.Moves, R)
		} else {
			val, err := strconv.Atoi(m)
			aoc22.CheckError(err)
			input.Moves = append(input.Moves, val)
		}
	}
	return input

}


func NextState(state State, board [][]Tile, n_steps int) State {
	pos := state.Pos
	dir := DIRECTIONS[state.Facing]
	n_rows := len(board)
	n_cols := len(board[0])

	for i := 0; i < n_steps; i++ {
		row := pos.row + dir.row
		col := pos.col + dir.col

		if dir.row == -1 && (row < 0 || board[row][col] == VOID) {
			row = n_rows - 1
			for board[row][col] == VOID {
				row--
			}
		}

		if dir.row == 1 && (row >= n_rows || board[row][col] == VOID) {
			row = 0
			for board[row][col] == VOID {
				row++
			}
		}

		if dir.col == -1 && (col < 0 || board[row][col] == VOID) {
			col = n_cols - 1
			for board[row][col] == VOID {
				col--
			}
		}

		if dir.col == 1 && (col >= n_cols || board[row][col] == VOID) {
			col = 0
			for board[row][col] == VOID {
				col++
			}
		}

		if board[row][col] == WALL {
			break
		}
		pos.row = row
		pos.col = col
	}
	return State{pos, state.Facing}
}

func CheckInside(inside bool) {
	if !inside {log.Panicf("Not a face")}
}

func NextPosCube(state State, board [][]Tile, n_steps int, cubeSize int) State {
	pos := state.Pos
	facing := state.Facing
	dir := DIRECTIONS[state.Facing]

	n_rows := len(board)
	n_cols := len(board[0])

	// harcoded only for my input
	neighbors := make(map[State]State)

	neighbors[State{Pos{0, 1}, UP}] = State{Pos{3, 0}, RIGHT}
	neighbors[State{Pos{0, 1}, LEFT}] = State{Pos{2, 0}, RIGHT}

	neighbors[State{Pos{0, 2}, RIGHT}] = State{Pos{2, 1}, LEFT}
	neighbors[State{Pos{0, 2}, UP}] = State{Pos{3, 0}, UP}
	neighbors[State{Pos{0, 2}, DOWN}] = State{Pos{1, 1}, LEFT}

	neighbors[State{Pos{1, 1}, RIGHT}] = State{Pos{0, 2}, UP}
	neighbors[State{Pos{1, 1}, LEFT}] = State{Pos{2, 0}, DOWN}

	neighbors[State{Pos{2, 0}, LEFT}] = State{Pos{0, 1}, RIGHT}
	neighbors[State{Pos{2, 0}, UP}] = State{Pos{1, 1}, RIGHT}

	neighbors[State{Pos{2, 1}, RIGHT}] = State{Pos{0, 2}, LEFT}
	neighbors[State{Pos{2, 1}, DOWN}] = State{Pos{3, 0}, LEFT}

	neighbors[State{Pos{3, 0}, RIGHT}] = State{Pos{2, 1}, UP}
	neighbors[State{Pos{3, 0}, DOWN}] = State{Pos{0, 2}, DOWN}
	neighbors[State{Pos{3, 0}, LEFT}] = State{Pos{0, 1}, DOWN}

	log.Printf("Going %s for %d steps", facing, n_steps)
	for i := 0; i < n_steps; i++ {
		face := Pos{pos.row / cubeSize, pos.col / cubeSize}
		dir = DIRECTIONS[facing]
		log.Printf("i=%d, pos=%v, face=%v, facing=%s", i, pos, face, facing)

		face_row := pos.row % cubeSize
		face_col := pos.col % cubeSize

		row := (pos.row + dir.row + n_rows) % n_rows
		col := (pos.col + dir.col + n_cols) % n_cols

		n, inside := neighbors[State{face, facing}]

		if facing == UP && board[row][col] == VOID {
			log.Printf("Switching %v -> %v, %s -> %s", face, n.Pos, facing, n.Facing)
			CheckInside(inside)
			switch n.Facing {
			case UP:
				row = n.Pos.row * cubeSize + cubeSize - 1
				col = n.Pos.col * cubeSize + face_col
			case RIGHT:
				row = n.Pos.row * cubeSize + face_col
				col = n.Pos.col * cubeSize
			default:
				log.Panicf("Not implemented: UP -> %s", n.Facing)
			}
			facing = n.Facing
		}

		if facing == DOWN && board[row][col] == VOID {
			log.Printf("Switching %v -> %v, %s -> %s", face, n.Pos, facing, n.Facing)
			CheckInside(inside)
			switch n.Facing {
			case LEFT:
				row = n.Pos.row * cubeSize + face_col
				col = n.Pos.col * cubeSize + cubeSize - 1
			case DOWN:
				row = n.Pos.row * cubeSize
				col = n.Pos.col * cubeSize + face_col
			default:
				log.Panicf("Not implemented: UP -> %s", n.Facing)
			}
			facing = n.Facing
		}

		if facing == RIGHT && board[row][col] == VOID {
			log.Printf("Switching %v -> %v, %s -> %s", face, n.Pos, facing, n.Facing)
			CheckInside(inside)
			switch n.Facing {
			case LEFT:
				row = n.Pos.row * cubeSize + (cubeSize - face_row - 1)
				col = n.Pos.col * cubeSize + cubeSize - 1
			case UP:
				row = n.Pos.row * cubeSize + cubeSize - 1
				col = n.Pos.col * cubeSize + face_row
			default:
				log.Panicf("Not implemented: UP -> %s", n.Facing)
			}
			facing = n.Facing
		}

		if facing == LEFT && board[row][col] == VOID {
			log.Printf("Switching %v -> %v, %s -> %s", face, n.Pos, facing, n.Facing)
			CheckInside(inside)
			switch n.Facing {
			case RIGHT:
				row = n.Pos.row * cubeSize + (cubeSize - face_row - 1)
				col = n.Pos.col * cubeSize
			case DOWN:
				row = n.Pos.row * cubeSize 
				col = n.Pos.col * cubeSize + face_row
			default:
				log.Panicf("Not implemented: UP -> %s", n.Facing)
			}
			facing = n.Facing
		}

		if board[row][col] == WALL {
			log.Printf("Hit wall at {%d, %d}", row, col)
			break
		}

		pos.row = row
		pos.col = col
	}
	return State{pos, facing}
}


func Walk(input Input, cubeSize int) int {
	col := 0
	for input.Board[0][col] != MOVABLE {
		col++
	}
	state := State{Pos{0, col}, RIGHT}

	for _, move := range input.Moves {
		switch m := move.(type) {
		case Turn:
			if m == L {
				state.Facing = (state.Facing + 3) % 4
			} else {
				state.Facing = (state.Facing + 5) % 4
			}
			log.Printf("New facing: %s", state.Facing)
		case int:
			if cubeSize > 0 {
				state = NextPosCube(state, input.Board, m, cubeSize)
			} else {
				state = NextState(state, input.Board, m)
			}
			log.Printf("New pos: %v, new facing: %s", state.Pos, state.Facing)
		}
	}

	return 1000 * (state.Pos.row + 1) + 4 * (state.Pos.col + 1) + int(state.Facing)
}

func PartOne(input Input) int {
	return Walk(input, 0)
}

func PartTwo(input Input, cubeSize int) int {
	return Walk(input, cubeSize)
}

func Day22() {
	log.Print("Getting Input")
	text, err := aoc22.GetInput(2022, 22)
	aoc22.CheckError(err)

	log.Print("Parsing Input")
	input := ParseInput(text)
	log.Printf("Done")

	log.Printf("Part 1: %d", PartOne(input))
	log.Printf("Part 2: %d", PartTwo(input, 50))
}
