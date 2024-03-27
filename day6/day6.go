package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := utils.ReadInputToStringArray("input.txt", "\r\n")
	if err != nil {
		log.Fatal(err)
	}
	raceTimes, distances, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	margin := getMargin(raceTimes, distances)
	fmt.Println(margin)
}

func parseInput(input []string) ([]int, []int, error) {
	timeString := strings.TrimPrefix(input[0], "Time:")
	raceTime, err := utils.ConvertArrayToInt(strings.Fields(timeString))
	if err != nil {
		return nil, nil, err
	}
	distanceString := strings.TrimPrefix(input[1], "Distance:")
	distance, err := utils.ConvertArrayToInt(strings.Fields(distanceString))
	if err != nil {
		return nil, nil, err
	}
	return raceTime, distance, nil
}

func getMargin(raceTimes, distances []int) int {
	margin := 1
	for i := range len(raceTimes) {
		margin *= getRecords(raceTimes[i], distances[i])
	}
	return margin
}

func getRecords(time, distance int) int {
	count := 0
	for t := range time {
		if (time-t)*t > distance {
			count++
		}
	}
	return count
}
