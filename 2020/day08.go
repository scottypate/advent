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
	filepath, err := filepath.Abs("2020/inputs/day08.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputVal, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(inputVal), "\n")
}

func getNumber(s string) int {
	re := regexp.MustCompile(`(\+|-)([0-9]*)`)
	operation := re.ReplaceAllString(s, "$1")
	number, _ := strconv.Atoi(re.ReplaceAllString(s, "$2"))

	if operation == "-" {
		return -number
	} else {
		return number
	}
}

// Execute the boot code
func executor(inputVal []string) (int, []int) {
	executionLog := make(map[int]int)
	var debugLog []int
	instructionIndex := 0
	acc := 0
	for i := 0; i < len(inputVal); i++ {
		executionLog[instructionIndex]++
		if executionLog[instructionIndex] > 1 {
			return acc, debugLog
		} else if instructionIndex == len(inputVal) {
			fmt.Println("Instructions completed exactly once.")
			return acc, debugLog
		}
		instruction := strings.Split(inputVal[instructionIndex], " ")
		increment := getNumber(instruction[1])

		if instruction[0] == "acc" {
			acc += increment
			debugLog = append(debugLog, instructionIndex)
			instructionIndex++
		} else if instruction[0] == "jmp" {
			debugLog = append(debugLog, instructionIndex)
			instructionIndex += increment
		} else if instruction[0] == "nop" {
			debugLog = append(debugLog, instructionIndex)
			instructionIndex++
		}
	}
	return 0, debugLog
}

// Find the nop or jmp operation that needs to be changed
func debugger(inputVal []string, debugLog []int) int {
	for _, element := range debugLog {
		modifiedInput := make([]string, len(inputVal))
		copy(modifiedInput, inputVal)
		instruction := strings.Split(modifiedInput[element], " ")
		if instruction[0] == "jmp" {
			modifiedInput[element] = strings.Replace(modifiedInput[element], "jmp", "nop", 1)
		} else if instruction[0] == "nop" {
			modifiedInput[element] = strings.Replace(modifiedInput[element], "nop", "jmp", 1)
		}
		acc, debugLog := executor(modifiedInput)
		if debugLog[len(debugLog)-1] == len(inputVal)-1 {
			return acc
		}
	}
	return 0
}

// Find the value of the accumulator when a line is executed twice.
func partOne(inputVal []string) int {
	acc, _ := executor(inputVal)
	return acc
}

// Fix the boot code to run exactly once.
func partTwo(inputVal []string) int {
	_, debugLog := executor(inputVal)
	acc := debugger(inputVal, debugLog)
	return acc
}

func main() {
	inputVal := getInput()
	fmt.Printf("Part One - The value of the Accumulator: %d\n", partOne(inputVal))
	fmt.Printf("Part Two - The value of the Accumulator: %d", partTwo(inputVal))
}
