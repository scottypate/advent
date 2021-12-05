package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinates struct {
	x, depth, aim int
}

// Read the input file.
func getInput() []string {
	file, err := os.Open("2021/inputs/day02.txt")

	if err != nil {
		log.Fatalf("Failed to open file!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var returnSlice []string
	for scanner.Scan() {
		lineStr := scanner.Text()
		returnSlice = append(returnSlice, lineStr)
	}
	return returnSlice
}

// Track horizontal position (x) and depth given a set of instructions
func partOne(inputVal []string) int {
	coordinates := Coordinates{0, 0, 0}
	var instruction []string
	for i := 0; i < len(inputVal); i++ {
		instruction = strings.Split(inputVal[i], " ")
		magnitude, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatal(err)
		}

		if instruction[0] == "forward" {
			coordinates.x += magnitude
		} else if instruction[0] == "down" {
			coordinates.depth += magnitude
		} else if instruction[0] == "up" {
			coordinates.depth -= magnitude
		}
	}
	return coordinates.x * coordinates.depth
}

// Track horizontal position (x), depth, and aim given a set of instructions
func partTwo(inputVal []string) int {
	coordinates := Coordinates{0, 0, 0}
	var instruction []string
	for i := 0; i < len(inputVal); i++ {
		instruction = strings.Split(inputVal[i], " ")
		magnitude, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatal(err)
		}

		if instruction[0] == "forward" {
			coordinates.x += magnitude
			coordinates.depth += coordinates.aim * magnitude
		} else if instruction[0] == "down" {
			coordinates.aim += magnitude
		} else if instruction[0] == "up" {
			coordinates.aim -= magnitude
		}
	}
	return coordinates.x * coordinates.depth
}

// Run both parts
func main() {
	inputVal := getInput()
	fmt.Printf("Part One answer is: %d \n", partOne(inputVal))
	fmt.Printf("Part Two answer is: %d \n", partTwo(inputVal))
}
