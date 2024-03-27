package main

import (
	"slices"
	"testing"
)

var (
	knownTimes          = []int{7, 15, 30}
	knownDistances      = []int{9, 40, 200}
	knownSingleTime     = 71530
	knownSingleDistance = 940200
)

func TestParseIput(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	raceTimes, distances, err := parseInputP1(input)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(raceTimes, knownTimes) {
		t.Fatalf("MISMATCH\n%v\n%v", knownTimes, raceTimes)
	}
	if !slices.Equal(distances, knownDistances) {
		t.Fatalf("MISMATCH\n%v\n%v", knownDistances, distances)
	}

	raceTime, distance, err := parseInputP2(input)
	if err != nil {
		t.Fatal(err)
	}
	if raceTime != knownSingleTime {
		t.Fatalf("MISMATCH\n%v\n%v", knownSingleTime, raceTime)
	}
	if distance != knownSingleDistance {
		t.Fatalf("MISMATCH\n%v\n%v", knownSingleDistance, distance)
	}

}

func TestMargin(t *testing.T) {
	knownMargin := 288
	margin := getMargin(knownTimes, knownDistances)
	if margin != knownMargin {
		t.Fatalf("MISMATCH\n%v\n%v", knownMargin, margin)
	}

	knownCount := 71503
	count := getRecords(knownSingleTime, knownSingleDistance)
	if count != knownCount {
		t.Fatalf("MISMATCH\n%v\n%v", knownCount, count)
	}
}
