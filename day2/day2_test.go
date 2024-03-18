package main

import (
	"fmt"
	"testing"
)

var gamesArray = []game{
	{[]gameSet{{4, 0, 3}, {1, 2, 6}, {0, 2, 0}}, 1},
	{[]gameSet{{0, 2, 1}, {1, 3, 4}, {0, 1, 1}}, 2},
	{[]gameSet{{20, 8, 6}, {4, 13, 5}, {1, 5, 0}}, 3},
	{[]gameSet{{3, 1, 6}, {6, 3, 0}, {14, 3, 15}}, 4},
	{[]gameSet{{6, 3, 1}, {1, 2, 2}}, 5},
}

func TestGamesParsing(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	games, err := parseGames(input)
	if err != nil {
		t.Fatal(err)
	}

	if fmt.Sprint(games) != fmt.Sprint(gamesArray) {
		t.Fatalf("was expecting %d but got %d", gamesArray, games)
	}
}

func TestPossibleGames(t *testing.T) {
	expectedGames := []game{
		{[]gameSet{{4, 0, 3}, {1, 2, 6}, {0, 2, 0}}, 1},
		{[]gameSet{{0, 2, 1}, {1, 3, 4}, {0, 1, 1}}, 2},
		{[]gameSet{{6, 3, 1}, {1, 2, 2}}, 5},
	}
	expectedSumIds := 8
	possibleGames, sumIds := getPossibleGames(gamesArray, 12, 13, 14)

	if fmt.Sprint(possibleGames) != fmt.Sprint(expectedGames) {
		t.Fatalf("MISMATCH\n%v\n%v", possibleGames, expectedGames)
	}

	if sumIds != expectedSumIds {
		t.Fatalf("MISMATCH\n%v\n%v", sumIds, expectedSumIds)
	}
}

func TestMinimumCubes(t *testing.T) {
	expectedGames := []game{
		{[]gameSet{{4, 2, 6}}, 1},
		{[]gameSet{{1, 3, 4}}, 2},
		{[]gameSet{{20, 13, 6}}, 3},
		{[]gameSet{{14, 3, 15}}, 4},
		{[]gameSet{{6, 3, 2}}, 5},
	}
	expectedPower := 2286
	minCubesRequiredPerGame, totalPower := getMinCubesRequired(gamesArray)

	if fmt.Sprint(minCubesRequiredPerGame) != fmt.Sprint(expectedGames) {
		t.Fatalf("MISMATCH\n%v\n%v", minCubesRequiredPerGame, expectedGames)
	}

	if totalPower != expectedPower {
		t.Fatalf("MISMATCH\n%v\n%v", totalPower, expectedPower)
	}
}