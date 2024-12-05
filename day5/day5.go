package day5

import (
	"fmt"
	"strings"

	"github.com/emahl/adventofcode2024/shared"
)

type Rule struct {
	Before int
	After  int
}

func Run() {
	rules, pageNumberUpdates := readRulesAndPageNumbersFromFile()
	part1(rules, pageNumberUpdates)
}

func part1(rules []Rule, pageNumberUpdates [][]int) {
	correctlyOrderedUpdates := getCorrectlyOrderedUpdates(pageNumberUpdates, rules)
	middleNumbers := getMiddleNumbers(correctlyOrderedUpdates)
	sum := 0

	for _, num := range middleNumbers {
		sum += num
	}
	fmt.Println("Sum of middle page numbers from correctly ordered updates:", sum)
}

func readRulesAndPageNumbersFromFile() ([]Rule, [][]int) {
	file, scanner := shared.ReadFile("day5/input.txt")
	defer file.Close()

	var rules []Rule
	var pageNumbers [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			rules = append(rules, Rule{
				Before: shared.ConvertToNumber(parts[0]),
				After:  shared.ConvertToNumber(parts[1]),
			})
		} else if strings.Contains(line, ",") {
			pageNumbers = append(pageNumbers, convertToNumbers(strings.Split(line, ",")))
		}
	}

	return rules, pageNumbers
}

func convertToNumbers(strs []string) []int {
	numbers := make([]int, len(strs))
	for i, s := range strs {
		numbers[i] = shared.ConvertToNumber(s)
	}
	return numbers
}

func getCorrectlyOrderedUpdates(pageNumberUpdates [][]int, rules []Rule) [][]int {
	var correctlyOrdered [][]int
	for _, update := range pageNumberUpdates {
		if isValidUpdate(update, rules) {
			correctlyOrdered = append(correctlyOrdered, update)
		}
	}
	return correctlyOrdered
}

func isValidUpdate(update []int, rules []Rule) bool {
	for i, pn := range update {
		if !isValidRule(pn, update[i:], rules) {
			return false
		}
	}
	return true
}

func isValidRule(pageNumber int, numbersAfter []int, rules []Rule) bool {
	for _, rule := range rules {
		if rule.After == pageNumber {
			for _, after := range numbersAfter {
				if rule.Before == after {
					return false
				}
			}
		}
	}
	return true
}

func getMiddleNumbers(correctlyOrderedUpdates [][]int) []int {
	var middleNumbers []int
	for _, update := range correctlyOrderedUpdates {
		middleNumbers = append(middleNumbers, update[(len(update)-1)/2])
	}
	return middleNumbers
}
