package elo

import (
	"fmt"
	"math"
)

var D float64 = 400.0
var K float64 = 32.0

/// Compute updated ratings for all players
func ComputeRatings(ratings []float64, ties []bool) ([]float64, error) {
	if len(ratings) != len(ties) {
		return nil, fmt.Errorf("ratings and equalities should have the same length")
	}

	racePositions := computePositions(ties)

	N := len(ratings)
	updatedRatings := make([]float64, N)
	for i := range racePositions {
		expected := computeExpectedScore(i, ratings)
		actual := computeActualScore(racePositions[i], N)
		updatedRating := computeUpdatedRating(ratings[i], expected, actual, N)
		updatedRating = round(updatedRating)
		updatedRatings[i] = updatedRating
	}
	return updatedRatings, nil
}

/// Round to closest integer
func round(rating float64) float64 {
	floored := float64(int(rating))
	if rating-floored >= 0.5 {
		return floored + 1
	} else {
		return floored
	}
}

// Computes a race position, base on ties
// For example, if 1st is tied with second, they will both have position 1.5 (instead of 1 and 2)
// This works for any amount of consecutive ties
func computePositions(ties []bool) []float64 {
	first := 0
	sum := 0
	out := make([]float64, len(ties))
	for i := range ties {
		sum += i + 1
		if !ties[i] {
			for k := first; k <= i; k++ {
				out[k] = float64(sum) / float64(1+i-first)
			}
			first = i + 1
			sum = 0
		}
	}
	return out
}

/// Returns the updated score for a player, given his expected and actual score
func computeUpdatedRating(rating, expectedScore, actualScore float64, N int) float64 {
	return rating + K*float64(N-1.0)*(actualScore-expectedScore)
}

/// Returns the actual score, given the player's position
/// 1 if the player finished 1st and 0 if he finished last
func computeActualScore(position float64, N int) float64 {
	return (float64(N) - position) / (float64(N*(N-1)) / 2.0)
}

/// Returns the expected score for a player, given all ratings
/// The expected score is between 0 and 1 and represents a probability of winning
func computeExpectedScore(currentIndex int, ratings []float64) float64 {
	N := len(ratings)
	sum := 0.0
	for i := range ratings {
		if i != currentIndex {
			sum += expectedTwoPlayersScore(ratings[currentIndex], ratings[i])
		}
	}
	return sum / (float64(N*(N-1)) / 2.0)
}

/// Returns the expected score for player A, given playerA and playerB's current ratings
func expectedTwoPlayersScore(ratingA, ratingB float64) float64 {
	return 1.0 / (1.0 + math.Pow(10.0, (ratingB-ratingA)/D))
}
