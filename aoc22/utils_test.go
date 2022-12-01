package aoc22

import "testing"

func TestSumInts(t *testing.T) {
	var result int

	result = Sum([]int{})
	if result != 0 {
		t.Errorf("Unexpected sum, got %d, expected 0", result)
	}

	result = Sum([]int{1, 2, 3})
	if result != 6 {
		t.Errorf("Unexpected sum, got %d, expected 6", result)
	}
}

func TestSumFloat64(t *testing.T) {
	var result float64

	result = Sum([]float64{})
	if result != 0 {
		t.Errorf("Unexpected sum, got %f, expected 0", result)
	}

	result = Sum([]float64{1.5, 2, 3})
	if result != 6.5 {
		t.Errorf("Unexpected sum, got %f, expected 6", result)
	}
}
