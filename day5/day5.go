package day5

import (
	"fmt"
	"sort"
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
	part2(rules, pageNumberUpdates)
}

func part1(rules []Rule, pageNumberUpdates [][]int) {
	correctlyOrderedUpdates := getOrderedUpdates(pageNumberUpdates, rules, true)
	fmt.Println("Sum of middle page numbers from correctly ordered updates:", sumOfMiddleNumbers(correctlyOrderedUpdates))
}

func part2(rules []Rule, pageNumberUpdates [][]int) {
	incorrectlyOrderedUpdates := getOrderedUpdates(pageNumberUpdates, rules, false)
	correctlyOrderedUpdates := orderUpdates(incorrectlyOrderedUpdates, rules)
	fmt.Println("Sum of middle page numbers from incorrectly ordered updates:", sumOfMiddleNumbers(correctlyOrderedUpdates))
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

func orderUpdates(incorrectlyOrderedUpdates [][]int, rules []Rule) [][]int {
	for _, update := range incorrectlyOrderedUpdates {
		sort.Slice(update, func(i, j int) bool {
			for _, rule := range rules {
				if rule.Before == update[i] && rule.After == update[j] {
					return true
				}
			}
			return false
		})
	}
	return incorrectlyOrderedUpdates
}

func getOrderedUpdates(pageNumberUpdates [][]int, rules []Rule, filterValid bool) [][]int {
	var orderedUpdates [][]int
	for _, update := range pageNumberUpdates {
		if (filterValid && isValidUpdate(update, rules)) || (!filterValid && !isValidUpdate(update, rules)) {
			orderedUpdates = append(orderedUpdates, update)
		}
	}
	return orderedUpdates
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

func sumOfMiddleNumbers(pageNumberUpdates [][]int) int {
	sum := 0
	for _, update := range pageNumberUpdates {
		middleIndex := (len(update) - 1) / 2
		sum += update[middleIndex]
	}
	return sum
}
