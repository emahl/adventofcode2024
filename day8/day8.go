package day8

import (
	"fmt"

	"github.com/emahl/adventofcode2024/shared"
)

func Run() {
	antennaMap := readAntennaMapFromFile()
	part1(antennaMap)
}

func readAntennaMapFromFile() [][]rune {
	file, scanner := shared.ReadFile("day8/input.txt")
	defer file.Close()

	var letterMatrix [][]rune
	for scanner.Scan() {
		letterMatrix = append(letterMatrix, []rune(scanner.Text()))
	}
	return letterMatrix
}

func part1(antennaMap [][]rune) {
	antennaPositions := getAllAntennaPositions(antennaMap)
	maxY, maxX := len(antennaMap)-1, len(antennaMap[0])-1
	antinodes := getAntinodes(antennaPositions, maxY, maxX)
	fmt.Println("Number of antinodes:", len(shared.GetUnique(antinodes)))
}

func getAllAntennaPositions(antennaMap [][]rune) map[rune][]shared.Position {
	var antennaIndices = make(map[rune][]shared.Position)
	for i := 0; i < len(antennaMap); i++ {
		for j := 0; j < len(antennaMap[i]); j++ {
			r := antennaMap[i][j]
			if r != '.' {
				antennaIndices[r] = append(antennaIndices[r], shared.Position{Y: i, X: j})
			}
		}
	}
	return antennaIndices
}

func getAntinodes(antennaPositions map[rune][]shared.Position, maxY, maxX int) []shared.Position {
	var antinodes []shared.Position
	for r := range antennaPositions {
		for _, p := range antennaPositions[r] {
			for _, p2 := range antennaPositions[r] {
				if p.X != p2.X && p.Y != p2.Y {
					antinodes = append(antinodes, shared.Position{X: p.X + (p.X - p2.X), Y: p.Y - (p2.Y - p.Y)})
				}
			}
		}
	}

	return filterValidPositions(antinodes, maxY, maxX)
}

func filterValidPositions(positions []shared.Position, maxY, maxX int) []shared.Position {
	var valid []shared.Position
	for _, p := range positions {
		if p.X >= 0 && p.Y >= 0 && p.X <= maxX && p.Y <= maxY {
			valid = append(valid, p)
		}
	}
	return valid
}
