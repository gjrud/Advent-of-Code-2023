package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

type part struct {
	num  int
	kind rune
	x, y int
}

type gear struct {
	num1, num2 int
	x, y       int
}

func main() {
	input, err := utils.ReadInputToStringArray("input.txt", "\r\n")
	if err != nil {
		log.Fatal(err)
	}

	parts, err := getPartNumbers(input)
	if err != nil {
		log.Fatal(err)
	}
	sum := 0
	for _, pn := range parts {
		sum += pn.num
	}
	fmt.Println(sum)

	gears := getGears(parts)
	sumGears := 0
	for _, g := range gears {
		sumGears += g.num1 * g.num2
	}
	fmt.Println(sumGears)
}

func getPartNumbers(input []string) ([]part, error) {
	partNums := make([]part, 0)
	for y, line := range input {
		for x := 0; x < len(line); x++ {
			v := rune(line[x])
			if unicode.IsDigit(v) {
				newPn, length, err := getNumbers(input, x, y)
				if err != nil {
					return nil, err
				}
				if newPn != (part{}) {
					partNums = append(partNums, newPn)
				}

				x += length - 1
			}
		}
	}
	return partNums, nil
}

func getNumbers(input []string, col, row int) (part, int, error) {
	strNum := ""
	for _, v := range input[row][col:] {
		if unicode.IsDigit(v) {
			strNum += string(v)
		} else {
			break
		}
	}

	yStart, yEnd := utils.Max(row-1, 0), utils.Min(row+2, len(input))
	xStart := utils.Max(col-1, 0)

	for y, line := range input[yStart:yEnd] {
		xEnd := utils.Min(col+len(strNum)+1, len(line))
		for x, v := range line[xStart:xEnd] {
			if !unicode.IsDigit(v) && v != '.' {
				num, err := strconv.Atoi(strNum)
				if err != nil {
					return part{}, 0, err
				}
				pn := part{num, v, x + xStart, y + yStart}
				return pn, len(strNum), nil
			}
		}
	}
	return part{}, len(strNum), nil
}

func getGears(parts []part) []gear {
	gears := make([]gear, 0)
	for i, p1 := range parts[:len(parts)-1] {
		if p1.kind != '*' {
			continue
		}

		g := gear{}
		for _, p2 := range parts[i+1:] {
			if p2.kind != '*' {
				continue
			}

			if p1.x == p2.x && p1.y == p2.y && g == (gear{}) {
				g = gear{p1.num, p2.num, p1.x, p2.y}
			} else if g != (gear{}) {
				break
			}
		}
		if g != (gear{}) {
			gears = append(gears, g)
		}
	}
	return gears
}
