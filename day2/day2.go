package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"strings"
)

type game struct {
	sets []gameSet
	id   int
}

type gameSet struct {
	red, green, blue int
}

func (g game) String() string {
	return fmt.Sprintf("\nGame : %d\n%v\n", g.id, g.sets)
}

func main() {
	input, err := utils.ReadInputToStringArray("input.txt", "\r\n")
	if err != nil {
		log.Fatal(err)
	}

	games, err := parseGames(input)
	if err != nil {
		log.Fatal(err)
	}

	sumIds := 0
	possibleGames, sumIds := getPossibleGames(games, 12, 13, 14)
	fmt.Println(possibleGames, sumIds)

	minCubesRequiredPerGame, totalPower := getMinCubesRequired(games)
	fmt.Println(minCubesRequiredPerGame, totalPower)
}

func parseGames(input []string) ([]game, error) {
	games := make([]game, 0)

	for i, l := range input {
		prefix := fmt.Sprintf("Game %d: ", i+1)
		cutLine := strings.TrimPrefix(l, prefix)
		setStrings := strings.Split(cutLine, ";")

		g, err := getGame(setStrings, i)
		if err != nil {
			return nil, err
		}

		games = append(games, g)
	}
	return games, nil
}

func getGame(setStrings []string, index int) (game, error) {
	f := func(c rune) bool {
		if c == ' ' || c == ',' {
			return true
		}
		return false
	}
	sets := make([]gameSet, 0)
	for _, setString := range setStrings {
		fields := strings.FieldsFunc(setString, f)
		set, err := parseSetFields(fields)
		if err != nil {
			return game{}, err
		}
		sets = append(sets, set)
	}
	return game{sets, index + 1}, nil
}

func parseSetFields(fields []string) (gameSet, error) {
	red, green, blue := 0, 0, 0
	for i := 0; i < len(fields); i = i + 2 {
		quantity, err := utils.ParseNumber(fields[i])
		if err != nil {
			return gameSet{}, err
		}

		switch fields[i+1] {
		case "red":
			red = quantity
		case "green":
			green = quantity
		case "blue":
			blue = quantity
		}
	}
	return gameSet{red, green, blue}, nil
}

func getPossibleGames(games []game, red, green, blue int) ([]game, int) {
	possibleGames := make([]game, 0)
	sum := 0

	for _, g := range games {
		possible := true
		for _, set := range g.sets {
			if set.red > red || set.green > green || set.blue > blue {
				possible = false
				break
			}
		}
		if possible {
			possibleGames = append(possibleGames, g)
			sum += g.id
		}
	}
	return possibleGames, sum
}

func getMinCubesRequired(games []game) ([]game, int) {
	minCubesRequiredPerGame := make([]game, 0)
	totalPower := 0

	for _, g := range games {
		minRed, minGreen, minBlue := 0, 0, 0
		for _, set := range g.sets {
			if set.red > minRed {
				minRed = set.red
			}
			if set.green > minGreen {
				minGreen = set.green
			}
			if set.blue > minBlue {
				minBlue = set.blue
			}
		}
		minCubesRequiredPerGame = append(minCubesRequiredPerGame, game{[]gameSet{{minRed, minGreen, minBlue}}, g.id})
		totalPower += minRed * minGreen * minBlue
	}
	return minCubesRequiredPerGame, totalPower
}
