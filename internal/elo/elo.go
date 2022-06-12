package elo

import (
	"math"
)

var D float64 = 400.0
var K float64 = 32.0

/// Compute updated ratings for all players
func ComputeRatings(ratings []float64) []float64 {
	updatedRatings := make([]float64, len(ratings))
	for i := range ratings {
		expected := computeExpectedScore(i, ratings)
		actual := computeActualScore(i+1, len(ratings))
		updatedRating := computeUpdatedRating(ratings[i], expected, actual, len(ratings))
		updatedRating = floorToDecimal(updatedRating)
		updatedRatings[i] = updatedRating
	}
	return updatedRatings
}

/// Returns the rating, but floored to the decimal. Ex: 10.1931 -> 10.1
func floorToDecimal(rating float64) float64 {
	return float64(int(rating*10)) / 10.0
}

/// Returns the updated score for a player, given his expected and actual score
func computeUpdatedRating(rating, expectedScore, actualScore float64, N int) float64 {
	return rating + K*float64(N-1.0)*(actualScore-expectedScore)
}

/// Returns the actual score, given the player's position
/// 1 if the player finished 1st and 0 if he finished last
func computeActualScore(position, N int) float64 {
	return float64(N-position) / (float64(N*(N-1)) / 2.0)
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
