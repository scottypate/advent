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
	filepath, err := filepath.Abs("2020/inputs/day09.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputVal, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(inputVal), "\n")
}

func findSums(preamble []string) []int {
	var results []int
	for _, i := range preamble {
		for _, j := range preamble {
			int1, _ := strconv.Atoi(i)
			int2, _ := strconv.Atoi(j)
			results = append(results, int1+int2)
		}
	}
	sort.Ints(results)
	return results
}

// Crack the encryption by figuring out which number isn't a sum
// of two of the first 25 numbers.
func partOne(inputVal []string) int {
	for i := 25; i < len(inputVal); i++ {
		preamble := inputVal[:25+i]
		possible := findSums(preamble)
		intMessage, _ := strconv.Atoi(inputVal[i])
		searchResult := sort.SearchInts(possible, intMessage)
		if possible[searchResult] != intMessage {
			return intMessage
		}
	}
	return 0
}

// Find contiguous numbers in the input that add up to the answer from partOne
func partTwo(inputVal []string, answer int) int {
	for i := 0; i < len(inputVal); i++ {
		int1, _ := strconv.Atoi(inputVal[i])
		total := int1
		vals := []int{int1}
		for ii := i + 1; ii < len(inputVal); ii++ {
			int2, _ := strconv.Atoi(inputVal[ii])
			vals = append(vals, int2)
			total += int2
			if total > answer {
				break
			} else if total == answer {
				sort.Ints(vals)
				min := vals[0]
				max := vals[len(vals)-1]
				return min + max
			} else if total < answer {
				continue
			}
		}
	}
	return 0
}

func main() {
	inputVal := getInput()
	partOneAnswer := partOne(inputVal)
	fmt.Printf("Part One - First non-compliant message: %d\n", partOneAnswer)
	fmt.Printf("Part Two - Encryption answer: %d", partTwo(inputVal, partOneAnswer))

}
