package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

func main() {
	input, err := utils.ReadInputToStringArray("input.txt", "\r\n")
	if err != nil {
		log.Fatal(err)
	}

	partNums, err := getPartNumbers(input)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, pn := range partNums {
		sum += pn
	}
	fmt.Println(sum)
}

func getPartNumbers(input []string) ([]int, error) {
	partNums := make([]int, 0)
	for y, line := range input {
		for x := 0; x < len(line); x++ {
			v := rune(line[x])
			if unicode.IsDigit(v) {
				newPartNum, length, err := getNumbers(input, x, y)
				if err != nil {
					return make([]int, 0), err
				}
				if newPartNum != 0 {
					partNums = append(partNums, newPartNum)
				}

				x += length - 1
			}
		}
	}
	return partNums, nil
}

func getNumbers(input []string, xStart, yStart int) (int, int, error) {
	strNum := ""
	for _, v := range input[yStart][xStart:] {
		if unicode.IsDigit(v) {
			strNum += string(v)
		} else {
			break
		}
	}

	for _, line := range input[utils.Max(yStart-1, 0):utils.Min(yStart+2, len(input))] {
		for _, v := range line[utils.Max(xStart-1, 0):utils.Min(xStart+len(strNum)+1, len(line))] {
			if !unicode.IsDigit(v) && v != '.' {
				partNum, err := strconv.Atoi(strNum)
				if err != nil {
					return 0, 0, err
				}
				return partNum, len(strNum), nil
			}
		}
	}
	return 0, len(strNum), nil
}
