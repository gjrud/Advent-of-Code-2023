package main

import (
	"log"
	"slices"
	"testing"
)

var expectedSeeds = []seed{
	{79, 81, 81, 81, 74, 78, 78, 82},
	{14, 14, 53, 49, 42, 42, 43, 43},
	{55, 57, 57, 53, 46, 82, 82, 86},
	{13, 13, 52, 41, 34, 34, 35, 35},
}

func TestParseInput(t *testing.T) {
	input := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}

	seeds, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}
	if !slices.Equal(expectedSeeds, seeds) {
		t.Fatalf("MISMATCH\n%v\n%v", expectedSeeds, seeds)
	}
}
