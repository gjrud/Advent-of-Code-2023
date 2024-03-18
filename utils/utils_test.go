package utils

import (
	"slices"
	"testing"
)

func TestReadInputToStringArray(t *testing.T) {
	expected := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	lines, err := ReadInputToStringArray("inputs\\sample_day1.txt", "\r\n")
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(lines, expected) {
		t.Fatalf("Was exepecting\n%v\nbut got instead\n%v", expected, lines)
	}
}
