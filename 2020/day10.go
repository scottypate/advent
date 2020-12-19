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
	filepath, err := filepath.Abs("2020/inputs/day10.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputVal, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(inputVal), "\n")
}

// Find the voltage differences given a list of adapters.
func partOne(inputVal []string) int {
	var voltages []int
	voltageGaps := make(map[int]int)

	for _, element := range inputVal {
		voltage, _ := strconv.Atoi(element)
		voltages = append(voltages, voltage)
	}
	sort.Ints(voltages)

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

func main() {
	inputVal := getInput()
	partOneAnswer := partOne(inputVal)
	fmt.Printf("Part One - voltage gap product: %d\n", partOneAnswer)
}
