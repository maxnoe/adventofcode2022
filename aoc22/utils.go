package aoc22

func Sum[V int | float64](values []V) V {
	var s V
	for _, val := range values {
		s += val
	}
	return s
}
