package elo

import (
	"math"
	"testing"

	"github.com/thalkz/kart/config"
)

var eloTestConfig = &config.Config{
	Elo: config.ConfigElo{
		D: 1000.0,
		K: 32.0,
	},
}

func TestEloRatings(T *testing.T) {
	var cfg = &config.Config{
		Elo: config.ConfigElo{
			D: 100.0,
			K: 32.0,
		},
	}

	in := []float64{
		1100.0,
		1000.0,
		1000.0,
		1000.0,
	}
	ties := []bool{
		false,
		false,
		false,
		false,
	}
	out, _ := ComputeRatings(cfg, in, ties)
	for i := range in {
		T.Logf("Player #%v: %v", i+1, out[i]-in[i])
	}
}

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

	newRatings, err := ComputeRatings(eloTestConfig, oldRatings, ties)
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
