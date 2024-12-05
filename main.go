package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/emahl/adventofcode2024/day1"
	"github.com/emahl/adventofcode2024/day2"
	"github.com/emahl/adventofcode2024/day3"
	"github.com/emahl/adventofcode2024/day4"
	"github.com/emahl/adventofcode2024/day5"
	"github.com/emahl/adventofcode2024/shared"
)

func main() {
	chosenDay := 0

	for chosenDay < 1 || chosenDay > 25 {
		chosenDay = gatherUserInput()
		if chosenDay < 1 || chosenDay > 25 {
			fmt.Println("Invalid day! Enter a number between 1 and 25")
		}
	}

	switch chosenDay {
	case 1:
		day1.Run()
	case 2:
		day2.Run()
	case 3:
		day3.Run()
	case 4:
		day4.Run()
	case 5:
		day5.Run()
	default:
		fmt.Println("Not implemented yet...")
	}
}

func gatherUserInput() int {
	fmt.Print("Choose day: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		return shared.ConvertToNumber(input)
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		shared.Check(err)
	}

	return -1
}
