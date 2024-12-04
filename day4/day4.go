package day4

import (
	"fmt"
	"strings"

	"github.com/emahl/adventofcode2024/shared"
)

// Idea: Recursive search.
// When start of word ('X') is found in matrix: Start 8 new word searches.
// A word search traverses the matrix either horizontally, vertically or diagonally, both backwards and forwards.
// If a word search finds the whole word "XMAS" it reports back true, if not false.

// Adjustments for part 2:
// Return coordinates and direction info for all found MASes in the matrix.
// Make sure to flip all MASes to be counted from the top and descend downard.
//
// Imagine the following MAS identified as "DiagonalUpRight" with coordinates SY: 2, SX: 0, EY: 0, EX: 2
//	|  S|
//	| A |
//	|M .|
// It will be "flipped" and identified as "DiagonalDownLeft" with SY: 0, SX: 2, EY: 2, EX: 0
//
// Using this information, all MASs placed in the DiagonalDownRight direction are processed and checked
// to see if they are part of an 'X-shape' by searching for a corresponding MAS placed in the DiagonalDownLeft direction.

const (
	Horizontal = iota
	HorizontalBackwards
	Vertical
	VerticalBackwards
	DiagonalDownLeft
	DiagonalDownRight
	DiagonalUpLeft
	DiagonalUpRight
)

type MASInfo struct {
	StartY    int
	StartX    int
	EndY      int
	EndX      int
	Direction int
}

func Run() {
	letterMatrix := readLetterMatrixFromFile()

	part1(letterMatrix)
	part2(letterMatrix)
}

func part1(letterMatrix [][]rune) {
	sum := searchForXmas(letterMatrix)
	fmt.Println("XMAS occurrences:", sum)
}

func part2(letterMatrix [][]rune) {
	sum := searchForXShapedMas(letterMatrix)
	fmt.Println("X-shaped MAS occurences:", sum)
}

func readLetterMatrixFromFile() [][]rune {
	var letterMatrix [][]rune
	file, scanner := shared.ReadFile("day4/input.txt")
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		letterMatrix = append(letterMatrix, runes)
	}

	return letterMatrix
}

func searchForXShapedMas(letterMatrix [][]rune) int {
	// Search for all MASes in the matrix and store coordinate info
	var masInfos []MASInfo
	for y := 0; y < len(letterMatrix); y++ {
		for x := 0; x < len(letterMatrix[y]); x++ {
			current := letterMatrix[y][x]
			if current == 'M' {
				masInfos = append(masInfos, startMasSearch(letterMatrix, y, x)...)
			}
		}
	}

	// Check how many MASes that are part of an X
	sum := 0
	for _, m := range masInfos {
		if m.Direction == DiagonalDownRight {
			if isPartOfX(m, masInfos) {
				sum++
			}
		}
	}
	return sum
}

func isPartOfX(masInfo MASInfo, allIndices []MASInfo) bool {
	for _, m := range allIndices {
		if m.Direction == DiagonalDownLeft && m.StartY == masInfo.StartY && m.EndX == masInfo.EndX-2 {
			return true
		}
	}

	return false
}

func searchForXmas(letterMatrix [][]rune) int {
	sum := 0

	for y := 0; y < len(letterMatrix); y++ {
		for x := 0; x < len(letterMatrix[y]); x++ {
			current := letterMatrix[y][x]
			if current == 'X' {
				sum += startSearch(letterMatrix, y, x)
			}
		}
	}

	return sum
}

func startMasSearch(letterMatrix [][]rune, y int, x int) []MASInfo {
	var masInfos []MASInfo
	diagonalDirections := []int{DiagonalDownLeft, DiagonalDownRight, DiagonalUpLeft, DiagonalUpRight}

	for _, d := range diagonalDirections {
		if found, endY, endX := search(letterMatrix, y, x, 'A', d, 2); found {
			var masInfo MASInfo
			if d == DiagonalDownLeft || d == DiagonalDownRight {
				masInfo = MASInfo{StartY: y, StartX: x, EndY: endY, EndX: endX, Direction: d}
			} else {
				// Flip coordinates to ensure all diagonals originate from the top and descend downward
				newDirection := DiagonalDownLeft
				if d == DiagonalUpLeft {
					newDirection = DiagonalDownRight
				}
				masInfo = MASInfo{StartY: endY, StartX: endX, EndY: y, EndX: x, Direction: newDirection}
			}
			masInfos = append(masInfos, masInfo)
		}
	}
	return masInfos
}

func startSearch(letterMatrix [][]rune, y int, x int) int {
	allDirections := []int{Horizontal, HorizontalBackwards, Vertical, VerticalBackwards, DiagonalDownLeft, DiagonalDownRight, DiagonalUpLeft, DiagonalUpRight}

	sum := 0
	for _, d := range allDirections {
		if found, _, _ := search(letterMatrix, y, x, 'M', d, 1); found {
			sum++
		}
	}
	return sum
}

func search(letterMatrix [][]rune, y int, x int, searchChar rune, searchDirection int, counter int) (bool, int, int) {
	y, x = moveIndices(y, x, searchDirection)
	if !validateIndices(letterMatrix, y, x) {
		return false, -1, -1
	}

	if letterMatrix[y][x] == searchChar {
		if counter == 3 {
			// We have found the word!
			return true, y, x
		}

		// Continue the search
		counter++
		return search(letterMatrix, y, x, nextSearchChar(searchChar), searchDirection, counter)
	}
	return false, -1, -1
}

func validateIndices(letterMatrix [][]rune, y int, x int) bool {
	if y < 0 || x < 0 {
		return false
	}
	if y >= len(letterMatrix) || x >= len(letterMatrix[y]) {
		return false
	}

	return true
}

func nextSearchChar(char rune) rune {
	xmas := "XMAS"

	index := strings.IndexRune(xmas, char)
	index++

	return []rune(xmas)[index]
}

func moveIndices(y int, x int, searchDirection int) (int, int) {
	offsets := map[int][2]int{
		Vertical:            {1, 0},
		VerticalBackwards:   {-1, 0},
		Horizontal:          {0, 1},
		HorizontalBackwards: {0, -1},
		DiagonalDownRight:   {1, 1},
		DiagonalDownLeft:    {1, -1},
		DiagonalUpRight:     {-1, 1},
		DiagonalUpLeft:      {-1, -1},
	}

	offset := offsets[searchDirection]
	y += offset[0]
	x += offset[1]

	return y, x
}
