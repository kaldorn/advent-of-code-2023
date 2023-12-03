package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestFindAndConcatFirstAndLastDigit(t *testing.T) {
	tests := []struct {
		given    string
		expected Game
	}{
		// Simple test to make sure things are parsed correctly.
		{
			given: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expected: Game{
				ID: 1,
				Rolls: []Roll{
					{
						Red:   4,
						Blue:  3,
						Green: 0,
					},
					{
						Red:   1,
						Green: 2,
						Blue:  6,
					},
					{
						Red:   0,
						Blue:  0,
						Green: 2,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		test := tc
		result := formatPlayedGames(test.given)
		// fmt.Println("String:", test.given, "Result:", result, "Expected:", test.expected)
		assert.DeepEqual(t, result, test.expected)
	}
}

func TestGetMaxCubesRolledPerColorPerGame(t *testing.T) {
	tests := []struct {
		given    string
		expected Roll
	}{
		// Simple test to make sure things are parsed correctly.
		{
			given: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expected: Roll{
				Red:   4,
				Blue:  6,
				Green: 2,
			},
		},
	}

	for _, tc := range tests {
		test := tc
		result := formatPlayedGames(test.given)
		// fmt.Println("String:", test.given, "Result:", result, "Expected:", test.expected)
		minimumCubesRequiredForGame := getMaxCubesRolledPerColorPerGame(result)
		assert.DeepEqual(t, minimumCubesRequiredForGame, test.expected)
	}
}
