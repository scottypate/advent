package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// Read the input file.
func getInput() []string {
	filepath, err := filepath.Abs("2020/inputs/day10test.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputVal, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(inputVal), "\n")
}

func convertList(inputVal []string) []int {
	var returnVal []int

	for _, element := range inputVal {
		convertedElement, err := strconv.Atoi(element)
		
		if err != nil {
			log.Fatal(err)
		}

		returnVal = append(returnVal, convertedElement)
	}
	sort.Ints(returnVal)

	return returnVal
}

// Find the voltage differences given a list of adapters.
func partOne(voltages []int) int {
	voltageGaps := make(map[int]int)

	for i, element := range voltages {
		if i == 0 {
			voltageGaps[voltages[i]]++
			continue
		}
		voltageGap := element - voltages[i-1]
		voltageGaps[voltageGap]++
	}
	// Multiple the 1 and 3 gap differences
	gapProduct := voltageGaps[1] * (voltageGaps[3] + 1)

	return gapProduct
}

// Find how many possible combinations of adapters can exist
func partTwo(voltages []int) int {

	allCombinations := 1

	for i, element := range voltages {
		nCombinations := 0
   		for ii := 1; ii < 4; ii++ {
   			if i + ii >= len(voltages) {
   				continue
   			}
   			diff := voltages[i + ii] - element
   			if diff <= 3 {
   				nCombinations++
   			}
   		}
   		allCombinations *= nCombinations
	}
	return allCombinations
}

func main() {
	inputVal := getInput()
	voltages := convertList(inputVal)
	partOneAnswer := partOne(voltages)
	partTwoAnswer := partTwo(voltages)
	fmt.Printf("Part One - voltage gap product: %d\n", partOneAnswer)
	fmt.Printf("Part Two - total combinations: %d\n", partTwoAnswer)
}
