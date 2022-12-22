package day21

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)

type Expression interface {
	String() string
}


type Op byte
const (
	Add Op = '+'
	Sub Op = '-'
	Mul Op = '*'
	Div Op = '/'
)

type Constant struct {
	Name string
	Value int
}


type BinaryOp struct {
	Name string
	Left string
	Op Op
	Right string
}


func (b BinaryOp) String() string  {
	return fmt.Sprintf("(%s %s %s)", b.Left, string(b.Op), b.Right)
}

func (c Constant) String() string  {
	if c.Name == "humn" {
		return fmt.Sprintf("humn=%d", c.Value)
	}
	return fmt.Sprintf("%d", c.Value)
}

var Inverse = map[Op]Op {
	Add: Sub,
	Sub: Add,
	Mul: Div,
	Div: Mul,
}

func ParseInput(input string) map[string]Expression {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	expressions := make(map[string]Expression)

	for _, line := range(lines) {
		parts := strings.Split(line, ": ")
		name := parts[0]
		expr := strings.Split(parts[1], " ")
		if len(expr) == 1 {
			val, err := strconv.Atoi(expr[0])
			aoc22.CheckError(err)
			expressions[name] = Constant{name, val}
		} else {
			expressions[name] = BinaryOp{name, expr[0], Op(expr[1][0]), expr[2]}
		}

	}
	return expressions

}

func CheckExists(exists bool, name string) {
	if !exists {
		log.Panicf("No expression named %s", name)
	}
}

func eval(name string, expressions map[string]Expression) int {
	expr, exists := expressions[name]
	CheckExists(exists, name)

	switch t := expr.(type) {
	case Constant:
		return t.Value
	case BinaryOp:
		left := eval(t.Left, expressions)
		right := eval(t.Right, expressions)

		switch t.Op {
		case Add: return left + right
		case Sub: return left - right
		case Mul: return left * right
		case Div: return left / right
		}
	default:
		log.Panicf("Unsupported expr: %v", t)
	}
	return -1
}

type FindResult int8
const (
	NOTFOUND FindResult = 0
	LEFT FindResult = 1
	RIGHT FindResult = 2
)

func Find(start string, target string, expressions map[string]Expression) FindResult {
	expr, exists := expressions[start]
	CheckExists(exists, start)

	switch t := expr.(type) {
	case BinaryOp:
		if t.Left == target {
			return LEFT
		}
		if t.Right == target {
			return RIGHT
		}

		if Find(t.Left, target, expressions) != NOTFOUND {
			return LEFT
		}

		if Find(t.Right, target, expressions) != NOTFOUND {
			return RIGHT
		}

		return NOTFOUND
	case Constant:
		return NOTFOUND
	}
	return NOTFOUND
}

func (r FindResult) String() string {
	switch r {
	case LEFT: return "left"
	case RIGHT: return "right"
	case NOTFOUND: return "no"
	default:
		panic("Uknown result")
	}
} 



func PartOne(expressions map[string]Expression) int {
	return eval("root", expressions)
}

func EvalHumn(humn int, e string, expressions map[string]Expression) int {
	before, _ := expressions["humn"]
	expressions["humn"] = Constant{"humn", humn}
	val := eval(e, expressions)
	expressions["humn"] = before
	return val
}

func PartTwo(expressions map[string]Expression) int {
	pos := Find("root", "humn", expressions)

	root, _ := expressions["root"]

	var expected_expr string
	var humn_expr string

	switch pos {
	case LEFT:
		expected_expr = root.(BinaryOp).Right
		humn_expr = root.(BinaryOp).Left
	case RIGHT:
		expected_expr = root.(BinaryOp).Left
		humn_expr = root.(BinaryOp).Right
	}

	expected := eval(expected_expr, expressions)

	humn := 0
	actual := 0
	low := 2
	high := math.MaxInt >> 7

	if EvalHumn(low, humn_expr, expressions) > EvalHumn(high, humn_expr, expressions) {
		low, high = high, low
	}

	for actual != expected { 
		high_val := EvalHumn(high, humn_expr, expressions)
		humn = (low + high) / 2 + 1
		actual = EvalHumn(humn, humn_expr, expressions)

		if actual < high_val && actual > expected {
			high = humn
		} else {
			low = humn
		}

	}

	return humn
}

func Day21() {
	log.Print("Getting Input")
	text, err := aoc22.GetInput(2022, 21)
	aoc22.CheckError(err)

	log.Print("Parsing Input")
	input := ParseInput(text)
	log.Printf("Done, %d expressions", len(input))

	log.Printf("Part 1: %d", PartOne(input))
	log.Printf("Part 2: %d", PartTwo(input))
}
