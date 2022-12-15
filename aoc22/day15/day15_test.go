package day15

import (
	"log"
	"reflect"
	"testing"
)

var test_input = `
Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
`

var test_sensors = []Sensor {
	{Pos{2, 18}, Pos{-2, 15}},
	{Pos{9, 16}, Pos{10, 16}},
	{Pos{13, 2}, Pos{15, 3}},
	{Pos{12, 14}, Pos{10, 16}},
	{Pos{10, 20}, Pos{10, 16}},
	{Pos{14, 17}, Pos{10, 16}},
	{Pos{8, 7}, Pos{2, 10}},
	{Pos{2, 0}, Pos{2, 10}},
	{Pos{0, 11}, Pos{2, 10}},
	{Pos{20, 14}, Pos{25, 17}},
	{Pos{17, 20}, Pos{21, 22}},
	{Pos{16, 7}, Pos{15, 3}},
	{Pos{14, 3}, Pos{15, 3}},
	{Pos{20, 1}, Pos{15, 3}},
}

func TestParse(t *testing.T) {
	sensors := ParseInput(test_input)
	for i, expected := range test_sensors {
		if !reflect.DeepEqual(sensors[i], expected) {
			t.Errorf("Parsing sensor %d failed, got:\n%v\nexpected:\n%v", i, sensors[i], expected)
		}
	}
}


func TestPartOne(t *testing.T) {
	answer := PartOne(test_sensors, 10)
	if answer != 26 {
		t.Errorf("Wrong answer: %d", answer)
	}
}

func TestPartTwo(t *testing.T) {
	answer := PartTwo(test_sensors, 20)
	log.Printf("Part 2: %d", answer)
	if answer != 56000011 {
		t.Errorf("Wrong answer: %d", answer)
	}
}
