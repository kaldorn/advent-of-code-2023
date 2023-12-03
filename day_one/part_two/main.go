package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	defer readFile.Close()

	numericStrings := []string{}
	for fileScanner.Scan() {
		numericStrings = append(numericStrings, findAndConcatFirstAndLastDigit(fileScanner.Text()))
	}

	sum := sumAllNumericStrings(numericStrings)

	fmt.Println(sum)
}

// Finds the first and last numeric characters in a string and returns that as a string.
// Additionally looks for spelled out digits and replaces each instance while still allowing for overlapping spelled digits.
func findAndConcatFirstAndLastDigit(str string) string {
	numbers := map[string]string{
		"nine":  "n9e",
		"eight": "e8t",
		"seven": "s7n",
		"six":   "s6x",
		"five":  "f5e",
		"four":  "f4r",
		"three": "t3e",
		"two":   "t2o",
		"one":   "o1e",
	}

	changedStr := str
	for key, val := range numbers {
		if strings.Contains(str, key) {
			changedStr = strings.ReplaceAll(changedStr, key, val)
		}
	}

	var first, last string
	end := len(changedStr) - 1
	for i, v := range changedStr {
		if v >= '0' && v <= '9' && first == "" {
			first = string(v)
		}

		if changedStr[end-i] >= '0' && changedStr[end-i] <= '9' && last == "" {
			last = string(changedStr[end-i])
		}

		if first != "" && last != "" {
			break
		}
	}
	return (first + last)
}

// Takes a slice of strings containing only numeric characters and adds them all together
func sumAllNumericStrings(strSlice []string) int {
	tmp := 0
	for _, v := range strSlice {
		res, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		tmp += res
	}
	return tmp
}
