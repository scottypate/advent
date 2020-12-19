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
	filepath, err := filepath.Abs("2020/inputs/day04.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputVal, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(inputVal), "\n\n")
}

func verifyKeys(passport map[string]string) bool {
	requiredKeys := []string{
		"byr", "iyr", "eyr", "hgt", "ecl", "hcl", "pid",
	}
	passportValid := true
	for _, key := range requiredKeys {
		_, ok := passport[key]
		if !ok {
			passportValid = false
			break
		}
	}
	return passportValid
}

func verifyMinMax(value int, min int, max int) bool {
	if value < min || value > max {
		return false
	} else {
		return true
	}
}

func verifyData(passport map[string]string) bool {
	// Verify byr
	byr, err := strconv.Atoi(passport["byr"])
	if err != nil || !verifyMinMax(byr, 1920, 2002) {
		return false
	}
	// Verify iyr
	iyr, err := strconv.Atoi(passport["iyr"])
	if err != nil || !verifyMinMax(iyr, 2010, 2020) {
		return false
	}
	// Verify eyr
	eyr, err := strconv.Atoi(passport["eyr"])
	if err != nil || !verifyMinMax(eyr, 2020, 2030) {
		return false
	}
	// Verify hgt
	hgtMatch, _ := regexp.MatchString("^[0-9]{1,3}(in|cm)", passport["hgt"])
	if !hgtMatch {
		return false
	}
	if strings.ContainsAny(passport["hgt"], "cm") {
		passport["hgt"] = strings.Replace(passport["hgt"], "cm", "", 1)
		hgt, _ := strconv.Atoi(passport["hgt"])
		if !verifyMinMax(hgt, 150, 193) {
			return false
		}
	} else if strings.ContainsAny(passport["hgt"], "in") {
		passport["hgt"] = strings.Replace(passport["hgt"], "in", "", 1)
		hgt, _ := strconv.Atoi(passport["hgt"])
		if !verifyMinMax(hgt, 59, 76) {
			return false
		}
	}
	// Verify hcl
	hclMatch, _ := regexp.MatchString("^#[0-9|a-f]{6}$", passport["hcl"])
	if !hclMatch {
		return false
	}
	// Verify ecl
	eclMatch, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", passport["ecl"])
	if !eclMatch {
		return false
	}
	// Verify pic
	picMatch, _ := regexp.MatchString("^[0-9]{9}$", passport["pid"])
	if !picMatch {
		return false
	}

	return true
}

// Count the number of passports that are have all required fields
func partOne(inputVal []string) (int, int) {
	valid := 0
	for _, element := range inputVal {
		passportMap := make(map[string]string)
		passport := strings.Fields(element)
		for _, field := range passport {
			valueSlice := strings.Split(string(field), ":")
			passportMap[valueSlice[0]] = valueSlice[1]
		}
		passportValid := verifyKeys(passportMap)
		if passportValid {
			valid++
		}
	}
	invalid := len(inputVal) - valid
	return valid, invalid
}

// Count the number of passports that are have all required fields AND correct data types
func partTwo(inputVal []string) (int, int) {
	valid := 0
	for _, element := range inputVal {
		passportMap := make(map[string]string)
		passport := strings.Fields(element)
		for _, field := range passport {
			valueSlice := strings.Split(string(field), ":")
			passportMap[valueSlice[0]] = valueSlice[1]
		}
		passportValid := verifyKeys(passportMap)
		if passportValid {
			passportValid = verifyData(passportMap)
		}
		if passportValid {
			valid++
		}
	}
	invalid := len(inputVal) - valid
	return valid, invalid
}

func main() {
	inputVal := getInput()
	partOneValid, partOneInvalid := partOne(inputVal)
	partTwoValid, partTwoInvalid := partTwo(inputVal)
	fmt.Printf("Part One - Valid: %d, Invalid: %d\n", partOneValid, partOneInvalid)
	fmt.Printf("Part Two - Valid: %d, Invalid: %d\n", partTwoValid, partTwoInvalid)
}
