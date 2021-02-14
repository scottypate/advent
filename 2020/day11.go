package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"strings"
)

type seatPosition struct {
	row, seat int
}

type seatAdjustment struct {
	rowChange, seatChange int
}

func adjustSeat(s seatPosition, a seatAdjustment) seatPosition {
	return seatPosition{
		s.row + a.rowChange,
		s.seat + a.seatChange,
	}
}

// Read the input file.
func getInput() []string {
	filepath, err := filepath.Abs("2020/inputs/day11.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputVal, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(inputVal), "\n")
}

// Return bool value for whether or not a seat in the direction of the initialAdjustment is occupied.
// If the adjacent parameter is true, then only examine the seat adjacent to the seatPosition.
func seatOccupancy(inputVal []string, seatPostition seatPosition, initialAdjustment seatAdjustment, adjacent bool) bool {
	nRows := len(inputVal)
	nSeats := len(inputVal[0])

	for i := 0; true; i++ {
		adjustment := seatAdjustment{
			initialAdjustment.rowChange + (initialAdjustment.rowChange * i),
			initialAdjustment.seatChange + (initialAdjustment.seatChange * i),
		}
		directionalSeat := adjustSeat(seatPostition, adjustment)
		//Break the loop if the directional adjustment is out of range of available seats.
		if directionalSeat.row >= nRows || directionalSeat.row < 0 || directionalSeat.seat >= nSeats || directionalSeat.seat < 0 {
			break
		}
		seatValue := inputVal[directionalSeat.row][directionalSeat.seat]
		if seatValue == '.' && adjacent {
			return false
		} else if seatValue == '.' && !adjacent {
			continue
		} else if seatValue == '#' {
			return true
		} else if seatValue == 'L' {
			return false
		}
	}
	return false
}

// Return the number of seats that are occupied
func countOccupiedSeats(inputVal []string, seat seatPosition, adjacent bool) int {
	nOccupied := 0
	adjustments := []int{-1, 0, 1}
	for _, rowChange := range adjustments {
		for _, seatChange := range adjustments {
			if rowChange == 0 && seatChange == 0 {
				continue
			}
			initialAdjustment := seatAdjustment{
				rowChange,
				seatChange,
			}
			seatStatus := seatOccupancy(
				inputVal, seat, initialAdjustment, adjacent,
			)
			if seatStatus {
				nOccupied++
			}
		}
	}
	return nOccupied
}

// Execute 1 round of seats being filled
func fillAdjacentSeats(inputVal []string) ([]string, int) {
	runningSum := 0
	filledSeats := make([]string, len(inputVal))
	for i, row := range inputVal {
		byteSlice := []byte(row)
		for ii, seat := range byteSlice {
			nOccupied := countOccupiedSeats(inputVal, seatPosition{i, ii}, true)
			if seat == 'L' && nOccupied == 0 {
				byteSlice[ii] = '#'
				runningSum++
			} else if seat == '#' && nOccupied >= 4 {
				byteSlice[ii] = 'L'
			} else if seat == '#' && nOccupied < 4 {
				runningSum++
			}
		}
		filledSeats[i] = string(byteSlice)
	}
	return filledSeats, runningSum
}

func fillDirectionalSeats(inputVal []string) ([]string, int) {
	runningSum := 0
	filledSeats := make([]string, len(inputVal))
	for i, row := range inputVal {
		byteSlice := []byte(row)
		for ii, seat := range byteSlice {
			nOccupied := countOccupiedSeats(inputVal, seatPosition{i, ii}, false)
			if seat == 'L' && nOccupied == 0 {
				byteSlice[ii] = '#'
				runningSum++
			} else if seat == '#' && nOccupied >= 5 {
				byteSlice[ii] = 'L'
			} else if seat == '#' && nOccupied < 5 {
				runningSum++
			}
		}
		filledSeats[i] = string(byteSlice)
	}
	return filledSeats, runningSum
}

// Apply a set of rules to a map of seats until no changes occur.
// When no further changes occur, return the number of Occupied seats (#).
func partOne(inputVal []string) int {
	var startingSeats []string
	changedSeats, nOccupied := fillAdjacentSeats(inputVal)
	for {
		if reflect.DeepEqual(changedSeats, startingSeats) {
			return nOccupied
		}
		startingSeats = changedSeats
		changedSeats, nOccupied = fillAdjacentSeats(changedSeats)
	}
	return 0
}

func partTwo(inputVal []string) int {
	var startingSeats []string
	changedSeats, nOccupied := fillDirectionalSeats(inputVal)
	for {
		if reflect.DeepEqual(changedSeats, startingSeats) {
			return nOccupied
		}
		startingSeats = changedSeats
		changedSeats, nOccupied = fillDirectionalSeats(changedSeats)
	}
	return 0
}

func main() {
	inputVal := getInput()
	partOneAnswer := partOne(inputVal)
	partTwoAnswer := partTwo(inputVal)
	fmt.Printf("Part One - Number of Occupied Seats: %d\n", partOneAnswer)
	fmt.Printf("Part Two - Number of Occupied Seats: %d\n", partTwoAnswer)
}
