package main

import (
	"testing"
)

func TestGetSumCalibrationValues(t *testing.T) {
	expected := 142
	input := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	sumCalibrationValues := getSumCalibrationValues(input)
	if sumCalibrationValues != expected {
		t.Fatalf("was expecting %d but got %d", expected, sumCalibrationValues)
	}
}

func TestGetSumCalibrationValues2(t *testing.T) {
	expected := 281
	input := []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	sumCalibrationValues := getSumCalibrationValues(input)
	if sumCalibrationValues != expected {
		t.Fatalf("was expecting %d but got %d", expected, sumCalibrationValues)
	}
}
