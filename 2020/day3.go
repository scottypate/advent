package main

import (
    "io/ioutil"
    "fmt"
    "log"
    "path/filepath"
    "strings"
    // "strconv"
)

type Coordinates struct {
    x, y int
}

// Read the input file.
func getInput() ([]string) {
    filepath, err := filepath.Abs("2020/inputs/day3.txt")
    if err != nil {
        log.Fatal(err)
    }

    inputVal, err := ioutil.ReadFile(filepath)
    if err != nil {
        log.Fatal(err)
    }

    return strings.Split(string(inputVal), "\n")
}

// Traverse the input at a slope of 3 right, 1 down and
// count the number of "trees"
func partOne(inputVal []string) (int) {
    position := Coordinates{0, 0}
    mapWidth := len(inputVal[0])
    nTrees := 0
    slope := []int{3, 1}
    
    for i, element := range inputVal {
        position.x = slope[0] * i
        position.y = slope[1] * i
        offset := (position.x/mapWidth) * mapWidth
        adjustedPosition := Coordinates{position.x - offset, position.y}
        if string(element[adjustedPosition.x]) == "#" {
            nTrees++
        }
    }

    return nTrees
}

// Traverse the input for a variety of slopes and multiple the results.
func partTwo(inputVal []string) (int) {
    position := Coordinates{0, 0}
    mapWidth := len(inputVal[0])
    var results []int
    slopes := [][]int{
        {1, 1},
        {3, 1},
        {5, 1}, 
        {7, 1},
        {1, 2},
    }
    
    for _, slope := range slopes {
        nTrees := 0
        for i := 0; i < len(inputVal); i += slope[1] {
            element := inputVal[i]
            position.x = slope[0] * i
            position.y = slope[1] * i
            offset := (position.x/mapWidth) * mapWidth
            adjustedPosition := Coordinates{position.x - offset, position.y}
            fmt.Println(i, offset, position, adjustedPosition)
            if string(element[adjustedPosition.x]) == "#" {
                nTrees++
            }
        }
        results = append(results, nTrees)
    }

    // Multiply the results together
    answer := 1

    for _, result := range results {
        answer *= result
    }

    return answer
}

func main() {
    inputVal := getInput()
    partOneAnswer := partOne(inputVal)
    partTwoAnswer := partTwo(inputVal)
    fmt.Printf("PartOne - Number of Trees: %d\n", partOneAnswer)
    fmt.Printf("PartTwo - Number of Trees Multiplied: %d", partTwoAnswer)
}
