package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"strings"
)

type card struct {
	winNums []int
	nums    []int
	copies  int
}

func (c card) winCount() int {
	winCount := 0
	for _, w := range c.winNums {
		for _, n := range c.nums {
			if w == n {
				winCount++
			}
		}
	}
	return winCount
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

	sumWorth := 0
	sumCopies := 0
	for i, c := range cards {
		winCount := c.winCount()
		for j := 1; j < len(cards) && j <= winCount; j++ {
			cards[i+j].copies += c.copies
		}
		if winCount > 0 {
			sumWorth += 1 << (winCount - 1)
		}
		sumCopies += c.copies
	}
	fmt.Println(sumWorth)
	fmt.Println(sumCopies)
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
		winNums, err := utils.ConvertArrayToInt(winStrings)
		if err != nil {
			return make([]card, 0), err
		}
		numStrings := strings.Fields(parts[1])
		nums, err := utils.ConvertArrayToInt(numStrings)
		if err != nil {
			return make([]card, 0), err
		}
		cards = append(cards, card{winNums, nums, 1})
	}
	return cards, nil
}
