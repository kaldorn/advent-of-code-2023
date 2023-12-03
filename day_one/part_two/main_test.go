package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestFindAndConcatFirstAndLastDigit(t *testing.T) {
	tests := []struct {
		given    string
		expected string
	}{
		{
			given:    "two1nine",
			expected: "29",
		},
		{
			given:    "eightwothree",
			expected: "83",
		},
		{
			given:    "abcone2threexyz",
			expected: "13",
		},
		{
			given:    "xtwone3four",
			expected: "24",
		},
		{
			given:    "4nineeightseven2",
			expected: "42",
		},
		{
			given:    "zoneight234",
			expected: "14",
		},
		{
			given:    "7pqrstsixteen",
			expected: "76",
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
