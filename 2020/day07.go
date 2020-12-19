package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Read the input file.
func getInput() []string {
	filepath, err := filepath.Abs("2020/inputs/day07.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputVal, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(inputVal), "\n")
}

// Turn the string input into a map structure where each outermost bag color is the key.
func parseInput(inputVal []string) map[string][]string {
	bagMap := make(map[string][]string)
	for _, element := range inputVal {
		bagSplit := strings.Split(string(element), " bags contain ")
		re := regexp.MustCompile(`([0-9]\s[a-z]*\s[a-z]*)\sbag[s]?(,?)[\.]?`)
		contents := re.ReplaceAllString(bagSplit[1], "$1$2")
		bagMap[bagSplit[0]] = strings.Split(contents, ", ")
	}
	return bagMap
}

// Create a map of the count of bags in each outermost bag
func parseCounts(bagMap map[string][]string) map[string]int {
	bagCounts := make(map[string]int)
	for key, value := range bagMap {
		for _, element := range value {
			re := regexp.MustCompile(`^([0-9])\s([a-z]*\s[a-z]*)`)
			replacement := re.ReplaceAllString(element, "$1")
			n, _ := strconv.Atoi(string(replacement))
			bagCounts[key] += n
		}
	}
	return bagCounts
}

// Given a bagColor return a slice of all the outermost bags which
// contain it.
func bagSearch(bagColor string, bagMap map[string][]string) []string {
	var matches []string
	for key, value := range bagMap {
		for _, element := range value {
			match, err := regexp.MatchString(bagColor, element)
			if match && err == nil {
				matches = append(matches, key)
			}
		}
	}
	return matches
}

// Create a map of the count of bags in each outermost bag
func bagCount(bagColor string, bagMap map[string][]string, bagCounts map[string]int) int {
	totalCount := bagCounts[bagColor]

	for _, match := range bagMap[bagColor] {
		re := regexp.MustCompile(`^([0-9])\s([a-z]*\s[a-z]*)`)
		multiple, _ := strconv.Atoi(re.ReplaceAllString(match, "$1"))
		color := re.ReplaceAllString(match, "$2")
		totalCount += bagCounts[color] * multiple
	}
	return totalCount
}

// Recursively search the bags
func recursiveSearch(bagColor string, bagMap map[string][]string) []string {
	totalMatches := bagSearch(bagColor, bagMap)

	for _, match := range totalMatches {
		totalMatches = append(totalMatches, recursiveSearch(match, bagMap)...)
	}
	return totalMatches
}

func recursiveCount(bagColor string, bagMap map[string][]string, bagCounts map[string]int) int {
	totalMatches := bagSearch(bagColor, bagMap)
	totalCount := bagCount(bagColor, bagMap, bagCounts)

	for _, match := range totalMatches {
		totalMatches = append(totalMatches, recursiveSearch(match, bagMap)...)
		totalCount += recursiveCount(match, bagMap, bagCounts)
	}
	return totalCount
}

// Count the number of outermost bags which can contain a shiny gold bag.
func partOne(inputVal []string) int {
	bagColor := "shiny gold"
	bagMap := parseInput(inputVal)
	numberMatches := make(map[string]int)
	matches := recursiveSearch(bagColor, bagMap)
	for _, match := range matches {
		numberMatches[match]++
	}
	return len(numberMatches)
}

// Count the number of bags a shiny gold bag must contain
func partTwo(inputVal []string) int {
	bagColor := "shiny gold"
	bagMap := parseInput(inputVal)
	fmt.Println(bagMap)
	bagCounts := parseCounts(bagMap)
	totalCount := recursiveCount(bagColor, bagMap, bagCounts)
	return totalCount
}

func main() {
	inputVal := getInput()
	fmt.Printf("Part One - Shiny Gold Bags: %d\n", partOne(inputVal))
	fmt.Printf("Part Two - Total Bags in Shiny Gold: %d", partTwo(inputVal))
}
