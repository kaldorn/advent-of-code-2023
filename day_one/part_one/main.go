package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
func findAndConcatFirstAndLastDigit(str string) string {
	var first, last string
	end := len(str) - 1
	for i, v := range str {
		if v >= '0' && v <= '9' && first == "" {
			first = string(v)
		}

		if str[end-i] >= '0' && str[end-i] <= '9' && last == "" {
			last = string(str[end-i])
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
