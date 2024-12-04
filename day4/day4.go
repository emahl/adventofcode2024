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

func Run() {
	letterMatrix := readLetterMatrixFromFile()

	part1(letterMatrix)
}

func part1(letterMatrix [][]rune) {
	sum := searchForXmas(letterMatrix)
	fmt.Println("XMAS occurrences:", sum)
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

func startSearch(letterMatrix [][]rune, y int, x int) int {
	allDirections := []int{Horizontal, HorizontalBackwards, Vertical, VerticalBackwards, DiagonalDownLeft, DiagonalDownRight, DiagonalUpLeft, DiagonalUpRight}

	sum := 0
	for _, d := range allDirections {
		if search(letterMatrix, y, x, 'M', d, 1) {
			sum++
		}
	}

	return sum
}

func search(letterMatrix [][]rune, y int, x int, searchChar rune, searchDirection int, counter int) bool {
	y, x = moveIndices(y, x, searchDirection)
	if !validateIndices(letterMatrix, y, x) {
		return false
	}

	if letterMatrix[y][x] == searchChar {
		if counter == 3 {
			// We have found XMAS!
			return true
		}

		// Continue the search
		counter++
		return search(letterMatrix, y, x, nextSearchChar(searchChar), searchDirection, counter)
	}
	return false
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
