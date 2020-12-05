package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

// Read the input file.
func getInput() []string {
	filepath, err := filepath.Abs("2020/inputs/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputVal, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(inputVal), "\n")
}

// Count the number of true results in a slice.
func nTrue(results []bool) int {
	n := 0
	for _, element := range results {
		if element {
			n++
		}
	}
	return n
}

// Password policy specifies occurrence thresholds (min, max)
// for the provided (letter).
func partOne(inputVal []string) (int, int) {
	valid := 0
	for _, element := range inputVal {
		// Parse the input string into its component parts.
		inputString := strings.Split(element, " ")
		minMax := strings.Split(inputString[0], "-")
		letter := strings.Replace(inputString[1], ":", "", 1)
		password := inputString[2]

		// Convert the min/max thresholds to integers
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])

		// Evaluate the password
		nOccurrences := strings.Count(password, letter)
		if nOccurrences >= min && nOccurrences <= max {
			valid++
		}
	}
	invalid := len(inputVal) - valid
	return valid, invalid
}

// Password policy specifies index positions (position1, position2)
// for the provided (letter).
func partTwo(inputVal []string) (int, int) {
	valid := 0
	for _, element := range inputVal {
		var results []bool
		// Parse the input string into its component parts.
		inputString := strings.Split(element, " ")
		indexPositions := strings.Split(inputString[0], "-")
		letter := strings.Replace(inputString[1], ":", "", 1)
		password := inputString[2]

		// Convert the required index positions to integers
		// the policy is indexed at 1 so we need to subtract 1.
		position1, _ := strconv.Atoi(indexPositions[0])
		position2, _ := strconv.Atoi(indexPositions[1])

		// Evaluate the password
		results = append(results, string(password[position1-1]) == letter)
		results = append(results, string(password[position2-1]) == letter)
		n := nTrue(results)

		if n == 1 {
			valid++
		}
	}
	invalid := len(inputVal) - valid
	return valid, invalid
}

func main() {
	inputVal := getInput()
	partOneValid, partOneInvalid := partOne(inputVal)
	partTwoValid, partTwoInvalid := partTwo(inputVal)
	fmt.Printf("PartOne - Valid: %d, Invalid: %d\n", partOneValid, partOneInvalid)
	fmt.Printf("PartTwo - Valid: %d, Invalid: %d\n", partTwoValid, partTwoInvalid)
}
