package shared

import (
	"bufio"
	"os"
	"strconv"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(fileName string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(fileName)
	Check(err)

	scanner := bufio.NewScanner(file)
	return file, scanner
}

func ConvertToNumber(numberStr string) int {
	convertedNumber, err := strconv.Atoi(numberStr)
	Check(err)

	return convertedNumber
}

func ConvertToString(number int) string {
	return strconv.Itoa(number)
}

func GetUnique(positions []Position) []Position {
	var uniquePositions []Position

	for _, p := range positions {
		if isUnique(uniquePositions, p) {
			uniquePositions = append(uniquePositions, p)
		}
	}

	return uniquePositions
}

func isUnique(uniquePositions []Position, p Position) bool {
	for _, u := range uniquePositions {
		if p.X == u.X && p.Y == u.Y {
			return false
		}
	}
	return true
}
