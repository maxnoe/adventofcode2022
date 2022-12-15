package day15

import (
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2022/aoc22"
)

type Pos struct {
	x int
	y int
}

type Sensor struct {
	pos Pos
	closest_beacon Pos
}

func ParseInput(input string) []Sensor {
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	sensors := make([]Sensor, len(lines))
	for i, line := range lines {
		groups := re.FindStringSubmatch(line)
		nums := make([]int, 4)
		for idx, group := range groups[1:] {
			val, err := strconv.Atoi(group)
			nums[idx] = val
			aoc22.CheckError(err)
		}
		sensors[i] = Sensor{Pos{nums[0], nums[1]}, Pos{nums[2], nums[3]}}
	}
	return sensors

}


func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}


func manhattan(p1, p2 Pos) int {
	return Abs(p1.x - p2.x) + Abs(p1.y - p2.y)
}


type Set map[Pos]bool


func PartOne(input []Sensor, row int) int {
	excluded := make(Set)
	beacons := make(Set)
	for _, sensor := range input {
		dist := manhattan(sensor.pos, sensor.closest_beacon)

		if sensor.closest_beacon.y == row {
			beacons[sensor.closest_beacon] = true
		}

		dy := Abs(sensor.pos.y - row)
		if dy > dist {
			continue
		}

		available := dist - dy + 1
		for dx := 0; dx < available; dx++ {
			excluded[Pos{sensor.pos.x + dx, row}] = true
			excluded[Pos{sensor.pos.x - dx, row}] = true
		}
	}
	return len(excluded) - len(beacons)
}

type Range struct {
	a int
	b int
}

type Ranges []Range

func (r Ranges) Less (i, j int) bool {
	if r[i].a == r[j].a {
		return r[i].b < r[j].b
	}
	return r[i].a < r[j].a
}

func (r Ranges) Swap (i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Ranges) Len() int {
	return len(r)
}

func PartTwo(sensors []Sensor, extent int) int {

	dists := make([]int, len(sensors))
	for i, sensor := range sensors {
		dists[i] = manhattan(sensor.pos, sensor.closest_beacon)
	}

	var beacon Pos

	excluded := make(Ranges, len(sensors))	
	for y := 0; y <= extent; y++ {

		n := 0
		for i, sensor := range sensors {
			dy := Abs(sensor.pos.y - y)
			remaining := dists[i] - dy
			if remaining < 1 {
				continue
			}

			excluded[n] = Range{
				sensor.pos.x - remaining,
				sensor.pos.x + remaining,
			}
			n++
		}
		sort.Sort(excluded[:n])
		x := 0
		for _, r := range excluded[:n] {
			if r.a <= x && r.b > x {
				x = r.b + 1
			}
		}
		if x <= extent {
			beacon = Pos{x, y}
			break
		}	
	}
	return 4000000 * beacon.x + beacon.y
}

func Day15() {
	log.Print("Getting Input")
	text, err := aoc22.GetInput(2022, 15)
	aoc22.CheckError(err)

	log.Print("Parsing Input")
	input := ParseInput(text)
	log.Printf("Done, %d sensors formations", len(input))

	log.Printf("Part 1: %d", PartOne(input, 2000000))
	log.Printf("Part 2: %d", PartTwo(input, 4000000))
}
