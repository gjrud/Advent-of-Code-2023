package main

import (
	"slices"
	"testing"
)

var (
	knownTimes     = []int{7, 15, 30}
	knownDistances = []int{9, 40, 200}
)

func TestParseIput(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	raceTimes, distances, err := parseInput(input)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(raceTimes, knownTimes) {
		t.Fatalf("MISMATCH\n%v\n%v", knownTimes, raceTimes)
	}
	if !slices.Equal(distances, knownDistances) {
		t.Fatalf("MISMATCH\n%v\n%v", knownDistances, distances)
	}
}

func TestMargin(t *testing.T) {
	knownMargin := 288
	margin := getMargin(knownTimes, knownDistances)
	if margin != knownMargin {
		t.Fatalf("MISMATCH\n%v\n%v", knownMargin, margin)
	}
}
