package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Read the input file.
func getInput() []int {
	file, err := os.Open("2021/inputs/day01.txt")

	if err != nil {
		log.Fatalf("Failed to open file!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var returnSlice []int
	for scanner.Scan() {
		lineStr := scanner.Text()
		num, _ := strconv.Atoi(lineStr)
		returnSlice = append(returnSlice, num)
	}
	return returnSlice
}

// Count the number of times the values increase
func partOne(inputVal []int) int {
	returnVal := 0
	for i := 1; i < len(inputVal); i++ {
		if inputVal[i] > inputVal[i-1] {
			returnVal++
		}
	}
	return returnVal
}

// Count the number of times the sum of a rolling window of three values increase
func partTwo(inputVal []int) int {
	returnVal := 0
	priorSum := 0
	for i := 1; i < len(inputVal)-2; i++ {
		sum := inputVal[i] + inputVal[i+1] + inputVal[i+2]
		if sum > priorSum {
			returnVal++
		}
		priorSum = sum
	}
	return returnVal
}

// Run both parts
func main() {
	inputVal := getInput()
	fmt.Printf("Part One answer is: %d \n", partOne(inputVal))
	fmt.Printf("Part Two answer is: %d", partTwo(inputVal))
}
