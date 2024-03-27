package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := utils.ReadInputToStringArray("input.txt", "\r\n")
	if err != nil {
		log.Fatal(err)
	}
	raceTimes, distances, err := parseInputP1(input)
	if err != nil {
		log.Fatal(err)
	}
	margin := getMargin(raceTimes, distances)
	fmt.Println(margin)

	raceTime, distance, err := parseInputP2(input)
	if err != nil {
		log.Fatal(err)
	}
	count := getRecords(raceTime, distance)
	fmt.Println(count)
}

func parseInputP1(input []string) ([]int, []int, error) {
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
func parseInputP2(input []string) (int, int, error) {
	timeString := strings.TrimPrefix(input[0], "Time:")
	raceTime, err := strconv.Atoi(strings.ReplaceAll(timeString, " ", ""))
	if err != nil {
		return 0, 0, err
	}
	distanceString := strings.TrimPrefix(input[1], "Distance:")
	distance, err := strconv.Atoi(strings.ReplaceAll(distanceString, " ", ""))
	if err != nil {
		return 0, 0, err
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
