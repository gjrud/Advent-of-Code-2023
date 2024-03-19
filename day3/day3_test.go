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

	parts, err := getPartNumbers(input)
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(parts, func(i, j int) bool {
		return parts[i].num < parts[j].num
	})

	partNums := make([]int, len(parts))
	for i, p := range parts {
		partNums[i] = p.num
	}

	if !slices.Equal(partNums, expectedSlice) {
		log.Fatalf("NOT EQUAL\n%v", parts)
	}

	sum := 0
	for _, pn := range partNums {
		sum += pn
	}
	if sum != expectedSum {
		log.Fatalf("NOT EQUAL\n%v", sum)
	}
}

func TestGetGears(t *testing.T) {
	parts := []part{
		{467, 42, 3, 1},
		{35, 42, 3, 1},
		{633, 35, 6, 3},
		{617, 42, 3, 4},
		{592, 43, 5, 5},
		{755, 42, 5, 8},
		{664, 36, 3, 8},
		{598, 42, 5, 8},
	}
	expectedSum := 467835

	gears := getGears(parts)

	sum := 0
	for _, g := range gears {
		sum += g.num1 * g.num2
	}

	if sum != expectedSum {
		log.Fatalf("NOT EQUAL\n%v", sum)
	}
}
