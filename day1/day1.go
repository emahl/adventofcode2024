package day1

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/emahl/adventofcode2024/shared"
)

func Run() {
	leftList, rightList := readListsFromFile()

	part1(leftList, rightList)
	part2(leftList, rightList)
}

func readListsFromFile() ([]int, []int) {
	file, scanner := shared.ReadFile("day1/input.txt")
	defer file.Close()

	var leftList []int
	var rightList []int

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")

		leftList = append(leftList, shared.ConvertToNumber(split[0]))
		rightList = append(rightList, shared.ConvertToNumber(split[1]))
	}

	return leftList, rightList
}

func part1(leftList []int, rightList []int) {
	sort.Ints(leftList)
	sort.Ints(rightList)

	distances := calculateDistances(leftList, rightList)
	sum := sumDistances(distances)

	fmt.Println("Sum distances:", sum)
}

func part2(leftList []int, rightList []int) {
	score := 0
	for _, num := range leftList {
		score += num * countNumberOfOccurences(num, rightList)
	}

	fmt.Println("Similarity score:", score)
}

func calculateDistances(list1 []int, list2 []int) []int {
	length := min(len(list1), len(list2))
	distances := make([]int, length)
	for i := 0; i < length; i++ {
		distances[i] = int(math.Abs(float64(list1[i] - list2[i])))
	}

	return distances
}

func sumDistances(array []int) int {
	sum := 0
	for _, num := range array {
		sum += num
	}
	return sum
}

func countNumberOfOccurences(number int, array []int) int {
	sum := 0
	for _, num := range array {
		if number == num {
			sum++
		}
	}
	return sum
}
