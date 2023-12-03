package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestFindAndConcatFirstAndLastDigit(t *testing.T) {
	/*
		Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
		Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
		Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
		Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
	*/
	tests := []struct {
		given    string
		expected bool
	}{
		{
			given:    "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expected: true,
		},
		{
			given:    "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expected: true,
		},
		{
			given:    "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expected: false,
		},
		{
			given:    "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			expected: false,
		},
		{
			given:    "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			expected: true,
		},
	}

	for _, tc := range tests {
		test := tc
		result := findAndConcatFirstAndLastDigit(test.given)
		// fmt.Println("String:", test.given, "Result:", result, "Expected:", test.expected)
		assert.Equal(t, result, test.expected, fmt.Sprintf("Result and expected are not equal. Result: %s Expected: %s String: %s", result, test.expected, test.given))
	}
}

func TestAddAllFirstAndLastDigits(t *testing.T) {
	test := struct {
		given    []string
		expected int
	}{
		given: []string{
			"12",
			"38",
			"15",
			"77",
		},
		expected: 142,
	}

	result := sumAllNumericStrings(test.given)
	assert.Equal(t, result, test.expected, "Result and expected are not equal")
}
