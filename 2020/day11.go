package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"reflect"
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

// Return the number of adjacent seats that are occupied.
func adjacentSeats(inputVal []string, seatPostition seatPosition) int {
	adjacentOccupied := 0
	adjustments := []int{-1, 0, 1}
	nRows := len(inputVal)
	nSeats := len(inputVal[0])
	for _, rowChange := range adjustments {
		for _, seatChange := range adjustments {
			adjacentSeat := adjustSeat(
				seatPostition, 
				seatAdjustment{rowChange, seatChange},
			)
			if adjacentSeat == seatPostition {
				continue
			} else if adjacentSeat.row >= nRows || adjacentSeat.row < 0 {
				continue
			} else if adjacentSeat.seat >= nSeats || adjacentSeat.seat < 0 {
				continue
			}
			seatStatus := inputVal[adjacentSeat.row][adjacentSeat.seat]
			if seatStatus == '#' {
				adjacentOccupied++
			}
		}
	}
	return adjacentOccupied
}

// Return bool value for whether or not a seat in the direction of the initialAdjustment is occupied.
func directionalSeatOccupied(inputVal []string, seatPostition seatPosition, initialAdjustment seatAdjustment) bool {
	nRows := len(inputVal)
	nSeats := len(inputVal[0])
	for i := 0; true; i++ {
		adjustment := seatAdjustment{
			initialAdjustment.rowChange + (initialAdjustment.rowChange * i),
			initialAdjustment.seatChange + (initialAdjustment.seatChange * i),
		}
		directionalSeat := adjustSeat(seatPostition, adjustment)

		if directionalSeat.row >= nRows || directionalSeat.row < 0 {
			break
		} else if directionalSeat.seat >= nSeats || directionalSeat.seat < 0 {
			break
		} else if inputVal[directionalSeat.row][directionalSeat.seat] == '.' {
			continue
		} else if inputVal[directionalSeat.row][directionalSeat.seat] == '#' {
			return true
		} else if inputVal[directionalSeat.row][directionalSeat.seat] == 'L' {
			return false
		}
	}
	return false
}

// Execute 1 round of seats being filled
func fillAdjacentSeats(inputVal []string) ([]string, int) {
	nOccupied := 0
	filledSeats := make([]string, len(inputVal))
	for i, row := range inputVal {
		byteSlice := []byte(row)
        	for ii, seat := range byteSlice {
        		if seat == '.' {
        			continue
        		}
        		adjacentOccupied := adjacentSeats(inputVal, seatPosition{i, ii})
        		if seat == 'L' && adjacentOccupied == 0 {
        			byteSlice[ii] = '#'
        			nOccupied++
        		} else if seat == '#' && adjacentOccupied >= 4 {
        			byteSlice[ii] = 'L'
        		} else if seat == '#' && adjacentOccupied < 4 {
        			nOccupied++
        		}
        	}
        	filledSeats[i] = string(byteSlice)
        }
        return filledSeats, nOccupied
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

// Apply a set of rules to a map of seats until no changes occur.
// When no further changes occur, return the number of Occupied seats (#).
func partTwo(inputVal []string) int {
	var startingSeats []string
	changedSeats, nOccupied := fillSeats(inputVal)
	for {

		if reflect.DeepEqual(changedSeats, startingSeats) {
			return nOccupied
		}
		startingSeats = changedSeats
		changedSeats, nOccupied = fillSeats(changedSeats)
	}
	return 0
}

func main() {
	inputVal := getInput()
	partOneAnswer := partOne(inputVal)
	fmt.Printf("Part One - Number of Occupied Seats: %d\n", partOneAnswer)
}
