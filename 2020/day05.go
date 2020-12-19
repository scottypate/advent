package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

// Read the input file.
func getInput() []string {
	filepath, err := filepath.Abs("2020/inputs/day05.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputVal, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(inputVal), "\n")
}

func makeSlice(min int, max int) []int {
	s := make([]int, max-min+1)
	for i := range s {
		s[i] = min + i
	}
	return s
}

func halveSlice(s []int, half string) []int {
	middle := len(s) / 2
	var returnVal []int

	if half == "lower" {
		returnVal = s[:middle]
	} else if half == "upper" {
		returnVal = s[middle:]
	}
	return returnVal
}

func calculateRow(instruction string) int {
	rows := makeSlice(0, 127)
	var half string
	for i := 0; i < 7; i++ {
		if string(instruction[i]) == "F" {
			half = "lower"
		} else if string(instruction[i]) == "B" {
			half = "upper"
		}
		rows = halveSlice(rows, half)
	}
	return rows[0]
}

func calculateSeat(instruction string) int {
	seats := makeSlice(0, 7)
	var half string
	for i := 7; i < 10; i++ {
		if string(instruction[i]) == "L" {
			half = "lower"
		} else if string(instruction[i]) == "R" {
			half = "upper"
		}
		seats = halveSlice(seats, half)
	}
	return seats[0]
}

// Figure out highest seat ID in the list of boarding passes
func partOne(inputVal []string) int {
	var seatIds []int
	for _, element := range inputVal {
		row := calculateRow(element)
		seat := calculateSeat(element)
		seatId := (row * 8) + seat
		seatIds = append(seatIds, seatId)
	}
	sort.Ints(seatIds)
	return seatIds[len(seatIds)-1]
}

// Find the missing seat ID in the list.
func partTwo(inputVal []string) int {
	var seatIds []int
	var mySeat int
	for _, element := range inputVal {
		row := calculateRow(element)
		seat := calculateSeat(element)
		seatId := (row * 8) + seat
		seatIds = append(seatIds, seatId)
	}
	sort.Ints(seatIds)
	min := seatIds[0]
	max := seatIds[len(seatIds)-1]

	for i := min; i < max; i++ {
		searchResult := sort.SearchInts(seatIds, i)
		if searchResult < len(seatIds) && seatIds[searchResult] == i {
			continue
		} else {
			mySeat = i
		}
	}

	return mySeat
}

func main() {
	inputVal := getInput()
	fmt.Printf("PartOne - Highest Seat ID: %d\n", partOne(inputVal))
	fmt.Printf("PartTwo - My Seat ID: %d", partTwo(inputVal))
}
