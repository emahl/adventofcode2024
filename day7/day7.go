package day7

import (
	"fmt"
	"strings"

	"github.com/emahl/adventofcode2024/shared"
)

type Equation struct {
	TestValue int
	Numbers   []int
}

func Run() {
	equations := readEquationsFromFile()

	part1(equations)
}

func readEquationsFromFile() []Equation {
	file, scanner := shared.ReadFile("day7/input.txt")
	defer file.Close()

	var equations []Equation
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ":")
		var numbers []int
		for _, number := range strings.Split(splits[1], " ") {
			if number != "" {
				numbers = append(numbers, shared.ConvertToNumber(number))
			}
		}
		equations = append(equations, Equation{TestValue: shared.ConvertToNumber(splits[0]), Numbers: numbers})
	}
	return equations
}

func part1(equations []Equation) {
	validEquations := getValidEquations(equations)
	sum := 0
	for _, equation := range validEquations {
		sum += equation.TestValue
	}
	fmt.Println("Total calibration result of valid equations:", sum)
}

func getValidEquations(equations []Equation) []Equation {
	var validEquations []Equation
	for _, equation := range equations {
		if isEquationValid(equation) {
			validEquations = append(validEquations, equation)
		}
	}
	return validEquations
}

func isEquationValid(equation Equation) bool {
	allSums := []int{}
	computeSums(equation.Numbers, 1, equation.Numbers[0], &allSums)

	for _, sum := range allSums {
		if equation.TestValue == sum {
			return true
		}
	}
	return false
}

func computeSums(numbers []int, index int, currentSum int, results *[]int) {
	if index == len(numbers) {
		*results = append(*results, currentSum)
		return
	}

	computeSums(numbers, index+1, currentSum+numbers[index], results)
	computeSums(numbers, index+1, currentSum*numbers[index], results)
}
