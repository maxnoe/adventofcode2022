package aoc22

import "log"

func Sum[V int | float64](values []V) V {
	var s V
	for _, val := range values {
		s += val
	}
	return s
}


func CheckError(err error) {
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
}


func AbsSign(x int) (int, int) {
	if x == 0 {
		return 0, 0
	}
	if x < 0 {
		return -x, -1
	}
	return x, 1
}
