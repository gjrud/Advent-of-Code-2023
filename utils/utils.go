package utils

import (
	"os"
	"strconv"
	"strings"
)

func ReadInputToStringArray(filepath string, sep string) ([]string, error) {
	output, err := readInput(filepath, sep)
	if err != nil {
		return make([]string, 0), err
	}
	return output, nil
}

func readInput(filepath string, sep string) ([]string, error) {
	f, err := os.ReadFile(filepath)
	if err != nil {
		return make([]string, 0), err
	}
	lines := strings.Split(string(f), sep)
	return lines, nil
}

func ConvertArrayToInt(lines []string) ([]int, error) {
	values := make([]int, 0)
	for _, l := range lines {
		v, err := ParseNumber(l)
		if err != nil {
			return make([]int, 0), err
		}
		values = append(values, v)
	}
	return values, nil
}

func ParseNumber(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
