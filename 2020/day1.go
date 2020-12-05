package main

import (
    "io/ioutil"
    "fmt"
    "log"
    "path/filepath"
    "strings"
    "strconv"
    "sort"
)

// Read the input file.
func getInput() ([]string) {
    filepath, err := filepath.Abs("2020/inputs/day1.txt")
    if err != nil {
        log.Fatal(err)
    }

    inputVal, err := ioutil.ReadFile(filepath)
    if err != nil {
        log.Fatal(err)
    }

    return strings.Split(string(inputVal), "\n")
}

// Convert the slice of strings to a slice of ints
func strToInt(slice []string) ([]int) {
    returnSlice := make([]int, len(slice))
    for i, element := range slice {
        element, err := strconv.Atoi(element)
        if err != nil {
            log.Fatal(err)
        }
        returnSlice[i] = element
    }
    sort.Ints(returnSlice)

    return returnSlice
}

// Find the 2 values in the slice that add up to the searchVal.
func partOne(intSlice []int, searchVal int) (int) {
    var returnVal int
    for _, element := range intSlice {
        complement := searchVal - element
        searchResult := sort.SearchInts(intSlice, complement)
        if searchResult < len(intSlice) && intSlice[searchResult] == complement {
            returnVal = element * complement
            break
        }
    }

    return returnVal
}

// Find the 3 values in the slice that add up to the searchVal.
func partTwo(intSlice []int, searchVal int) (int) {
    var returnVal int
    for i1, element1 := range intSlice {
        for i2, element2 := range intSlice {
            if i1 == i2 {
                continue
            }
            complement := searchVal - (element1 + element2)
            searchResult := sort.SearchInts(intSlice, complement)
            if searchResult < len(intSlice) && intSlice[searchResult] == complement {
                returnVal = element1 * element2 * complement
                break
            }
        }
    }

    return returnVal
}

// Run both parts
func main() {
    searchVal := 2020
    inputVal := getInput()
    intSlice := strToInt(inputVal)
    fmt.Println(partOne(intSlice, searchVal))
    fmt.Println(partTwo(intSlice, searchVal))
}