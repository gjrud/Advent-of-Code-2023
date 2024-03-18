package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"strings"
	"unicode"
)

func main() {
	input, err := utils.ReadInputToStringArray("input.txt", "\r\n")
	if err != nil {
		log.Fatalf("Encoundered Error:\n%v", err)
	}

	sumCalibrationValues := getSumCalibrationValues(input)
	fmt.Println(sumCalibrationValues)
}

func getSumCalibrationValues(lines []string) int {
	sum := 0
	for _, l := range lines {
		digits := getDigits(l)
		sum += digits[0]*10 + digits[len(digits)-1]
	}
	return sum
}

type numerals struct {
	name  string
	value int
}

func getDigits(s string) []int {
	digits := make([]int, len(s))
	words := []numerals{
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
	}
	for i, v := range s {
		if unicode.IsDigit(v) {
			digits[i] = int(v - '0')
		} else {
			digits[i] = -1
		}
	}
	for _, w := range words {
		index := len(s)
		for {
			match := strings.LastIndex(s[:index], w.name)
			if match == -1 {
				break
			}
			index = match
			digits[index] = w.value
		}
	}
	return cleanDigits(digits)
}

func cleanDigits(digits []int) []int {
	i := 0
	for _, v := range digits {
		if v != -1 {
			digits[i] = v
			i++
		}
	}
	return digits[:i]
}
