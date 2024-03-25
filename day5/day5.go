package main

import (
	"aoc2023/utils"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type seed struct {
	num   int
	soil  int
	fert  int
	water int
	light int
	temp  int
	humi  int
	loc   int
}

type mapRange struct {
	destStart int
	origStart int
	length    int
}

func main() {
	input, err := utils.ReadInputToStringArray("input.txt", "\r\n")
	if err != nil {
		log.Fatal(err)
	}

	seeds, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i].loc < seeds[j].loc
	})
	fmt.Println(seeds[0].loc)
}

func parseInput(input []string) ([]seed, error) {
	seeds := make([]seed, 0)
	seedNums, err := getSeedNums(input[0])
	if err != nil {
		return make([]seed, 0), err
	}

	allMaps := make([][]mapRange, 7)
	pos := 3
	for i := 1; i < 8; i++ {
		mapRanges, err := getMapRanges(input[pos:])
		if err != nil {
			return make([]seed, 0), err
		}
		allMaps[i-1] = mapRanges
		pos += len(mapRanges) + 2
	}

	for _, seedNum := range seedNums {
		res := make([]int, 8)
		res[0] = seedNum
		for i, mapRanges := range allMaps {
			res[i+1] = evaluateRanges(res[i], mapRanges)
		}
		seeds = append(seeds, seed{
			res[0],
			res[1],
			res[2],
			res[3],
			res[4],
			res[5],
			res[6],
			res[7],
		})
	}

	return seeds, nil
}

func getSeedNums(s string) ([]int, error) {
	seedNums := make([]int, 0)
	s = strings.TrimPrefix(s, "seeds:")
	for _, f := range strings.Fields(s) {
		num, err := strconv.Atoi(f)
		if err != nil {
			return make([]int, 0), err
		}
		seedNums = append(seedNums, num)
	}
	return seedNums, nil
}

func getMapRanges(s []string) ([]mapRange, error) {
	rangeMaps := make([]mapRange, 0)
	for _, l := range s {
		if l == "" {
			break
		}
		fieldStrings := strings.Fields(l)
		if len(fieldStrings) != 3 {
			return make([]mapRange, 0), errors.New("too many fields for rangemap")
		}
		fieldInts := make([]int, len(fieldStrings))
		for i, v := range fieldStrings {
			num, err := strconv.Atoi(v)
			if err != nil {
				return make([]mapRange, 0), err
			}
			fieldInts[i] = num
		}
		rangeMaps = append(rangeMaps, mapRange{fieldInts[0], fieldInts[1], fieldInts[2]})
	}
	// sort.Slice(rangeMaps, func(i, j int) bool {
	// 	return rangeMaps[i].origStart > rangeMaps[j].origStart
	// })
	return rangeMaps, nil
}

func evaluateRanges(input int, mapRanges []mapRange) int {
	result := input
	for _, mr := range mapRanges {
		if input >= mr.origStart && input < mr.origStart+mr.length {
			return mr.destStart + input - mr.origStart
		}
	}
	return result
}
