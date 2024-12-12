package day8

import (
	"fmt"

	"github.com/emahl/adventofcode2024/shared"
)

func Run() {
	antennaMap := readAntennaMapFromFile()
	antennaPositions := getAllAntennaPositions(antennaMap)
	maxY, maxX := len(antennaMap)-1, len(antennaMap[0])-1

	part1(antennaPositions, maxY, maxX)
	part2(antennaPositions, maxY, maxX)
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

func part1(antennaPositions map[rune][]shared.Position, maxY, maxX int) {
	antinodes := getAntinodes(antennaPositions, maxY, maxX, false)
	fmt.Println("Number of antinodes:", len(shared.GetUnique(antinodes)))
}

func part2(antennaPositions map[rune][]shared.Position, maxY, maxX int) {
	antinodes := getAntinodes(antennaPositions, maxY, maxX, true)
	fmt.Println("Number of antinodes with resonant harmonics:", len(shared.GetUnique(antinodes)))
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

func getAntinodes(antennaPositions map[rune][]shared.Position, maxY, maxX int, includeResonantHarmonics bool) []shared.Position {
	var antinodes []shared.Position
	for r := range antennaPositions {
		for _, p := range antennaPositions[r] {
			for _, p2 := range antennaPositions[r] {
				if p.X != p2.X && p.Y != p2.Y {
					counter := 1
					for {
						position := shared.Position{X: p.X - ((p2.X - p.X) * counter), Y: p.Y - ((p2.Y - p.Y) * counter)}
						if !isValidPosition(position, maxY, maxX) {
							break
						}
						antinodes = append(antinodes, position)
						if !includeResonantHarmonics {
							break // Stop after one iteration if we are not counting resonant harmonics
						}
						counter++
					}

					// Add the original position if resonant harmonics are included
					if includeResonantHarmonics {
						antinodes = append(antinodes, p)
					}
				}
			}
		}
	}

	return antinodes
}

func isValidPosition(p shared.Position, maxY, maxX int) bool {
	return p.X >= 0 && p.Y >= 0 && p.X <= maxX && p.Y <= maxY
}
