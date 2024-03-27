package main

import (
	"slices"
	"testing"
)

var (
	expectedCards = []card{
		{[]int{41, 48, 83, 86, 17}, []int{83, 86, 6, 31, 17, 9, 48, 53}, 1},
		{[]int{13, 32, 20, 16, 61}, []int{61, 30, 68, 82, 17, 32, 24, 19}, 1},
		{[]int{1, 21, 53, 59, 44}, []int{69, 82, 63, 72, 16, 21, 14, 1}, 1},
		{[]int{41, 92, 73, 84, 69}, []int{59, 84, 76, 51, 58, 5, 54, 83}, 1},
		{[]int{87, 83, 26, 28, 32}, []int{88, 30, 70, 12, 93, 22, 82, 36}, 1},
		{[]int{31, 18, 13, 56, 72}, []int{74, 77, 10, 23, 35, 67, 36, 11}, 1},
	}

	isCardEqual = func(e1, e2 card) bool {
		equal := slices.Equal(e1.winNums, e2.winNums)
		if equal {
			equal = slices.Equal(e1.nums, e2.nums)
		}
		return equal
	}
)

func TestParseInput(t *testing.T) {
	input := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	cards, err := parseInput(input)
	if err != nil {
		t.Fatal(err)
	}

	if !slices.EqualFunc(cards, expectedCards, isCardEqual) {
		t.Fatalf("MISMATCH\n%v\n%v", expectedCards, cards)
	}
}

func TestWorth(t *testing.T) {
	expectedWorth := 13

	sum := 0
	for _, c := range expectedCards {
		wincount := c.winCount()
		if wincount > 0 {
			sum += 1 << (wincount - 1)
		}
	}
	if sum != expectedWorth {
		t.Fatalf("MISMATCH\n%d\n%d", expectedWorth, sum)
	}
}

func TestCopies(t *testing.T) {
	cards := make([]card, len(expectedCards))
	copy(cards, expectedCards)
	expectedCopies := 30

	sum := 0
	for i, c := range cards {
		winCount := c.winCount()
		for j := 1; j < len(cards) && j <= winCount; j++ {
			cards[i+j].copies += c.copies
		}
		sum += c.copies
	}
	if sum != expectedCopies {
		t.Fatalf("MISMATCH\n%d\n%d", expectedCopies, sum)
	}
}
