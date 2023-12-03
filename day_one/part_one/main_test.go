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
			given:    "1abc2",
			expected: "12",
		},
		{
			given:    "pqr3stu8vwx",
			expected: "38",
		},
		{
			given:    "a1b2c3d4e5f",
			expected: "15",
		},
		{
			given:    "treb7uchet",
			expected: "77",
		},
		{
			given:    "8qntwokkffvkrlgoneightfnm",
			expected: "88",
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
