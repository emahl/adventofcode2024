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
