package day3

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/emahl/adventofcode2024/shared"
)

type MulInstruction struct {
	X int
	Y int
}

func Run() {
	memory := readMemoryFromFile()

	part1(memory)
	part2(memory)
}

func readMemoryFromFile() string {
	file, scanner := shared.ReadFile("day3/input.txt")
	defer file.Close()

	lines := ""
	for scanner.Scan() {
		lines += scanner.Text()
	}
	return lines
}

func part1(memory string) {
	mulInstructions := getMulInstructions(memory)
	sum := sumMulInstructions(mulInstructions)

	fmt.Println("Mul instruction sum:", sum)
}

func part2(memory string) {
	mulInstructions := getEnabledMulInstructions(memory)
	sum := sumMulInstructions(mulInstructions)

	fmt.Println("Enabled mul instruction sum:", sum)
}

var mulInstructionRegexp = `mul\([\d]+,[\d]+\)`

func getMulInstructions(memory string) []MulInstruction {
	var instructions []MulInstruction
	re := regexp.MustCompile(mulInstructionRegexp)
	matches := re.FindAllString(memory, -1)

	for _, m := range matches {
		instructions = append(instructions, parseMulInstruction(m))
	}

	return instructions
}

func getEnabledMulInstructions(memory string) []MulInstruction {
	var instructions []MulInstruction

	mre := regexp.MustCompile(mulInstructionRegexp)
	mulIndices := mre.FindAllStringIndex(memory, -1)

	dore := regexp.MustCompile(`do\(\)`)
	doIndices := dore.FindAllStringIndex(memory, -1)

	dontre := regexp.MustCompile(`don\'t\(\)`)
	dontIndices := dontre.FindAllStringIndex(memory, -1)

	allIndices := mergeAndSortIndices(mulIndices, doIndices, dontIndices)

	isEnabled := true
	for _, indices := range allIndices {
		if startIndexExists(doIndices, indices[0]) {
			isEnabled = true
		} else if startIndexExists(dontIndices, indices[0]) {
			isEnabled = false
		} else if isEnabled {
			instructions = append(instructions, parseMulInstruction(memory[indices[0]:indices[1]]))
		}
	}

	return instructions
}

func mergeAndSortIndices(x [][]int, y [][]int, z [][]int) [][]int {
	var allIndices [][]int
	allIndices = append(allIndices, x...)
	allIndices = append(allIndices, y...)
	allIndices = append(allIndices, z...)

	sort.Slice(allIndices, func(i, j int) bool {
		return allIndices[i][0] < allIndices[j][0]
	})

	return allIndices
}

func parseMulInstruction(s string) MulInstruction {
	s = strings.Trim(s, "mul(")
	s = strings.Trim(s, ")")

	split := strings.Split(s, ",")
	return MulInstruction{X: shared.ConvertToNumber(split[0]), Y: shared.ConvertToNumber(split[1])}
}

func sumMulInstructions(mulInstructions []MulInstruction) int {
	sum := 0
	for _, m := range mulInstructions {
		sum += m.X * m.Y
	}
	return sum
}

func startIndexExists(indicesSlice [][]int, startIndex int) bool {
	for _, item := range indicesSlice {
		if item[0] == startIndex {
			return true
		}
	}
	return false
}
