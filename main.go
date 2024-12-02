package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/emahl/adventofcode2024/day1"
	"github.com/emahl/adventofcode2024/day2"
	"github.com/emahl/adventofcode2024/shared"
)

func main() {
	chosenDay := gatherUserInput()
	if chosenDay < 1 || chosenDay > 25 {
		fmt.Println("Invalid day! Retry...")
		return
	}

	switch chosenDay {
	case 1:
		day1.Run()
	case 2:
		day2.Run()
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
