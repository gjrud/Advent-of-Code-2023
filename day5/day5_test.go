package main

import (
	"log"
	"slices"
	"testing"
)

var (
	expectedSeeds    = []int{79, 14, 55, 13}
	expectedLocation = 35
	expectedMapFunc  = mapFunction{
		[]mapPiece{{22, 0, 13},
			{29, 14, 14},
			{21, 15, 21},
			{68, 22, 25},
			{-25, 26, 43},
			{17, 44, 49},
			{-30, 50, 51},
			{-8, 52, 53},
			{31, 54, 58},
			{35, 59, 61},
			{-6, 62, 65},
			{31, 66, 68},
			{4, 69, 69},
			{-70, 70, 70},
			{3, 71, 81},
			{-36, 82, 91},
			{-32, 92, 92},
			{-25, 93, 97},
			{-31, 98, 98},
			{-80, 99, 99},
			{0, 100, 9223372036854775807},
		},
	}
)

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

	seeds, mapFunc, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	if !slices.Equal(seeds, expectedSeeds) {
		t.Fatalf("MISMATCH\n%v\n%v", expectedSeeds, seeds)
	}
	if !slices.Equal(mapFunc.mapPieces, expectedMapFunc.mapPieces) {
		t.Fatalf("MISMATCH\n%v\n%v", expectedMapFunc, mapFunc)
	}
}
