package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type card struct {
	winNums []int
	nums    []int
}

func (c card) worth() int {
	winCount := 0
	for _, w := range c.winNums {
		for _, n := range c.nums {
			if w == n {
				winCount++
			}
		}
	}

	if winCount == 0 {
		return 0
	} else {
		return 1 << (winCount - 1)
	}
}

func main() {
	input, err := utils.ReadInputToStringArray("input.txt", "\r\n")
	if err != nil {
		log.Fatal(err)
	}

	cards, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, c := range cards {
		sum += c.worth()
	}
	fmt.Println(sum)
}

func parseInput(input []string) ([]card, error) {
	cards := make([]card, 0)
	for _, line := range input {
		line = line[strings.IndexRune(line, ':')+1:]
		parts := strings.FieldsFunc(line, func(r rune) bool {
			if r == '|' {
				return true
			} else {
				return false
			}
		})
		winStrings := strings.Fields(parts[0])
		winNums, err := getNums(winStrings)
		if err != nil {
			return make([]card, 0), err
		}
		numStrings := strings.Fields(parts[1])
		nums, err := getNums(numStrings)
		if err != nil {
			return make([]card, 0), err
		}
		cards = append(cards, card{winNums, nums})
	}
	return cards, nil
}

func getNums(s []string) ([]int, error) {
	nums := make([]int, len(s))
	for j, ws := range s {
		num, err := strconv.Atoi(ws)
		if err != nil {
			return make([]int, 0), err
		}
		nums[j] = num
	}
	return nums, nil
}
