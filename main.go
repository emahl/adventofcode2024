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
	"github.com/emahl/adventofcode2024/day6"
	"github.com/emahl/adventofcode2024/day7"
	"github.com/emahl/adventofcode2024/day8"
	"github.com/emahl/adventofcode2024/shared"
)

var days = map[int]func(){
	1: day1.Run,
	2: day2.Run,
	3: day3.Run,
	4: day4.Run,
	5: day5.Run,
	6: day6.Run,
	7: day7.Run,
	8: day8.Run,
}

func main() {
	chosenDay := 0

	for chosenDay < 1 || chosenDay > 25 {
		chosenDay = gatherUserInput()
		if chosenDay < 1 || chosenDay > 25 {
			fmt.Println("Invalid day! Enter a number between 1 and 25")
		}
	}

	if runner, exists := days[chosenDay]; exists {
		runner()
	} else {
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
