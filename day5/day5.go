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

type mapFunction struct {
	mapPieces []mapPiece
}

func (mf *mapFunction) fillGaps() {
	filled := make([]mapPiece, 0)
	sort.Slice(mf.mapPieces, func(i, j int) bool {
		return mf.mapPieces[i].start < mf.mapPieces[j].start
	})

	if mf.mapPieces[0].start != 0 {
		filled = append(filled, mapPiece{
			shift: 0,
			start: 0,
			end:   mf.mapPieces[0].start - 1,
		})
	}
	filled = append(filled, mf.mapPieces[0])

	for i := 1; i < len(mf.mapPieces); i++ {
		prev := mf.mapPieces[i-1]
		gap := mf.mapPieces[i].start - prev.end
		if gap > 1 {
			filled = append(filled, mapPiece{
				shift: 0,
				start: prev.end + 1,
				end:   mf.mapPieces[i].start - 1,
			})
		}
		filled = append(filled, mf.mapPieces[i])
	}

	last := filled[len(filled)-1]
	if last.end != int(^uint(0)>>1) {
		filled = append(filled, mapPiece{
			shift: 0,
			start: filled[len(filled)-1].end + 1,
			end:   int(^uint(0) >> 1),
		})
	}

	mf.mapPieces = filled
}

type mapPiece struct {
	shift int
	start int
	end   int
}

func main() {
	input, err := utils.ReadInputToStringArray("input.txt", "\r\n")
	if err != nil {
		log.Fatal(err)
	}

	seeds, mapFunc, err := parseInput(input)
	if err != nil {
		log.Fatal(err)
	}

	locationP1 := getBestLocationP1(seeds, mapFunc)
	fmt.Println(locationP1)

	locationP2 := getBestLocationP2(seeds, mapFunc)
	fmt.Println(locationP2)
}

func parseInput(input []string) ([]int, mapFunction, error) {
	seedNums, err := getSeedNums(input[0])
	if err != nil {
		return nil, mapFunction{}, err
	}

	mapFunctions := make([]mapFunction, 7)
	pos := 3
	for i := 1; i < 8; i++ {
		mf, err := getMapFunction(input[pos:])
		if err != nil {
			return nil, mapFunction{}, err
		}
		mapFunctions[i-1] = mf
		pos += len(mf.mapPieces) + 2
	}
	mapFunc := mergeMapFunctions(mapFunctions)

	return seedNums, mapFunc, nil
}

func getSeedNums(s string) ([]int, error) {
	s = strings.TrimPrefix(s, "seeds:")
	seedNums, err := utils.ConvertArrayToInt(strings.Fields(s))
	if err != nil {
		return nil, err
	}
	return seedNums, nil
}

func getMapFunction(s []string) (mapFunction, error) {
	pieces := make([]mapPiece, 0)
	for _, l := range s {
		if l == "" {
			break
		}
		fieldStrings := strings.Fields(l)
		if len(fieldStrings) != 3 {
			return mapFunction{}, errors.New("too many fields for rangemap")
		}
		fieldInts := make([]int, len(fieldStrings))
		for i, v := range fieldStrings {
			num, err := strconv.Atoi(v)
			if err != nil {
				return mapFunction{}, err
			}
			fieldInts[i] = num
		}
		pieces = append(pieces, mapPiece{
			shift: fieldInts[0] - fieldInts[1],
			start: fieldInts[1],
			end:   fieldInts[1] + fieldInts[2] - 1,
		})
	}
	return mapFunction{pieces}, nil
}

func mergeMapFunctions(mf []mapFunction) mapFunction {
	result := mf[0]
	result.fillGaps()
	for i := 1; i < len(mf); i++ {
		pieces := make([]mapPiece, 0)
		mf[i].fillGaps()
		for _, v1 := range result.mapPieces {
			for _, v2 := range mf[i].mapPieces {
				merged := mergePieces(v1, v2)
				if merged != (mapPiece{}) {
					pieces = append(pieces, merged)
				}
			}
		}
		result.mapPieces = pieces
	}

	return result
}

func mergePieces(a, b mapPiece) mapPiece {
	bLow := b.start - a.shift
	bUp := b.end - a.shift
	if b.end == int(^uint(0)>>1) && a.shift < 0 {
		bUp = b.end
	}
	if a.end < bLow || bUp < a.start {
		return mapPiece{}
	}

	start := utils.Max(a.start, bLow)
	end := utils.Min(a.end, bUp)
	shift := a.shift + b.shift
	return mapPiece{
		shift: shift,
		start: start,
		end:   end,
	}
}

func getBestLocationP1(seeds []int, mapFunc mapFunction) int {
	bestLocation := int(^uint(0) >> 1)
	for _, seed := range seeds {
		for _, piece := range mapFunc.mapPieces {
			if seed < piece.start || seed > piece.end {
				continue
			}
			newLocation := seed + piece.shift
			if bestLocation > newLocation {
				bestLocation = newLocation
			}
		}
	}
	return bestLocation
}

func getBestLocationP2(seeds []int, mapFunc mapFunction) int {
	bestLocation := int(^uint(0) >> 1)
	for i := 0; i < len(seeds); i += 2 {
		for _, piece := range mapFunc.mapPieces {
			if seeds[i] > piece.end || seeds[i]+seeds[i+1]-1 < piece.start {
				continue
			}
			newLocation := utils.Max(piece.start, seeds[i]) + piece.shift
			if bestLocation > newLocation {
				bestLocation = newLocation
			}
		}
	}
	return bestLocation
}
