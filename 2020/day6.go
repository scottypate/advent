package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	// "regexp"
	// "strconv"
	"strings"
)

// Read the input file.
func getInput() []string {
	filepath, err := filepath.Abs("2020/inputs/day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputVal, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(inputVal), "\n\n")
}

// Count and sum the number of "yes" answers to the questions.
func partOne(inputVal []string) int {
	results := 0
	for _, element := range inputVal {
		answers := make(map[string]int)
		element = strings.ReplaceAll(element, "\n", "")
		for _, answer := range element {
			answers[string(answer)] = strings.Count(element, string(answer))
		}
		results += len(answers)
	}
	return results

}

// Count and sum the number of questions which everyone in the group answered.
func partTwo(inputVal []string) int {
	results := 0
	for _, element := range inputVal {
		answers := make(map[string]int)
		nIndividuals := len(strings.Split(string(element), "\n"))
		element := strings.ReplaceAll(element, "\n", "")
		for _, answer := range element {
			nOccurences := strings.Count(element, string(answer))
			if nOccurences == nIndividuals {
				answers[string(answer)] = strings.Count(element, string(answer))
			}
		}
		results += len(answers)
	}
	return results

}

func main() {
	inputVal := getInput()
	fmt.Printf("Part One - Sum of yes questions: %d\n", partOne(inputVal))
	fmt.Printf("Part Tow - Sum of yes questions for everyone: %d", partTwo(inputVal))
}
