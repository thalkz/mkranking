package elo

import (
	"math"
	"testing"
)

func TestComputeRatings(T *testing.T) {
	oldRatings := []float64{
		1000.0,
		1000.0,
		1000.0,
		1000.0,
		1000.0,
	}
	ties := []bool{
		false,
		true,
		true,
		false,
		false,
	}

	newRatings, err := ComputeRatings(oldRatings, ties)
	if err != nil {
		T.Errorf("failed to compute ratings: %v", err)
	}

	oldTotal := sum(oldRatings)
	newTotal := sum(newRatings)
	if !equals(oldTotal, newTotal) {
		T.Errorf("total actual score should be %v, got %v", oldTotal, newTotal)
	}
}

func equals(actual, expected float64) bool {
	return math.Abs(actual-expected) < 5
}

func sum(arr []float64) float64 {
	sum := 0.0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}
