package main

import (
	"log"
	"slices"
	"sort"
	"testing"
)

func TestGetPartNumbers(t *testing.T) {
	input := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	expectedSlice := []int{35, 467, 592, 598, 617, 633, 664, 755}
	expectedSum := 4361

	partNumbers, err := getPartNumbers(input)
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(partNumbers, func(i, j int) bool {
		return partNumbers[i] < partNumbers[j]
	})

	if !slices.Equal(partNumbers, expectedSlice) {
		log.Fatalf("NOT EQUAL\n%v", partNumbers)
	}

	sum := 0
	for _, pn := range partNumbers {
		sum += pn
	}
	if sum != expectedSum {
		log.Fatalf("NOT EQUAL\n%v", sum)
	}
}
