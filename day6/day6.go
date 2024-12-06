package day6

import (
	"fmt"

	"github.com/emahl/adventofcode2024/shared"
)

const (
	Up = iota
	Right
	Down
	Left
)

type Position struct {
	X int
	Y int
}

type GuardPosition struct {
	Position
	Direction int
}

func Run() {
	mappedArea, guardStartingPosition := readMapFromFile()

	part1(mappedArea, guardStartingPosition)
	part2(mappedArea, guardStartingPosition)
}

func readMapFromFile() ([][]rune, GuardPosition) {
	file, scanner := shared.ReadFile("day6/input.txt")
	defer file.Close()

	var mappedArea [][]rune
	var guardPosition GuardPosition
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		mappedArea = append(mappedArea, runes)
		for x, r := range runes {
			if r == '^' {
				guardPosition = GuardPosition{Position: Position{X: x, Y: y}, Direction: Up}
			}
		}
		y++
	}

	// Remove the guard position from the map
	mappedArea[guardPosition.Y][guardPosition.X] = '.'

	return mappedArea, guardPosition
}

func part1(mappedArea [][]rune, guardStartingPosition GuardPosition) {
	visitedPositions := getVisitedPositions(mappedArea, guardStartingPosition)
	sum := len(getUnique(visitedPositions))
	fmt.Println("Number of distinct guard positions:", sum)
}

func part2(mappedArea [][]rune, guardStartingPosition GuardPosition) {
	obstructionPositions := getObstructionsCreatingLoops(mappedArea, guardStartingPosition)
	sum := len(obstructionPositions)
	fmt.Println("Number of obstruction positions:", sum)
}

func getVisitedPositions(mappedArea [][]rune, guardStartingPosition GuardPosition) []GuardPosition {
	visitedPositionsLookup := make(map[GuardPosition]bool)
	currentPosition := guardStartingPosition

	for {
		newGuardPosition := moveGuard(mappedArea, currentPosition)

		// Stop if the guard moves outside the mapped area
		if newGuardPosition.X == -1 && newGuardPosition.Y == -1 {
			break
		}

		// If revisiting a position, return an empty slice
		if visitedPositionsLookup[newGuardPosition] {
			return []GuardPosition{}
		}

		visitedPositionsLookup[newGuardPosition] = true
		currentPosition = newGuardPosition
	}

	result := make([]GuardPosition, 0, len(visitedPositionsLookup))
	for p := range visitedPositionsLookup {
		result = append(result, p)
	}
	return result
}

func getObstructionsCreatingLoops(mappedArea [][]rune, guardStartingPosition GuardPosition) []Position {
	var obstructionPositions []Position

	for y := 0; y < len(mappedArea); y++ {
		for x := 0; x < len(mappedArea[y]); x++ {
			// Skip non-empty positions and the guard's starting position
			if mappedArea[y][x] != '.' || (y == guardStartingPosition.Y && x == guardStartingPosition.X) {
				continue
			}

			// Create a copy of the mapped area to test with the new obstruction
			areaCopy := make([][]rune, len(mappedArea))
			for i := range mappedArea {
				areaCopy[i] = append([]rune{}, mappedArea[i]...)
			}

			// Add new obstruction
			areaCopy[y][x] = '#'
			if len(getVisitedPositions(areaCopy, guardStartingPosition)) == 0 {
				obstructionPositions = append(obstructionPositions, Position{X: x, Y: y})
			}
		}
	}

	return obstructionPositions
}

func moveGuard(area [][]rune, currentGuardPosition GuardPosition) GuardPosition {
	nextX := -1
	nextY := -1

	for {
		nextX, nextY = getNextPosition(currentGuardPosition)

		if nextX < 0 || nextY < 0 || nextY >= len(area) || nextX >= len(area[0]) {
			nextX = -1
			nextY = -1
			break
		}

		if area[nextY][nextX] == '#' {
			switch currentGuardPosition.Direction {
			case Up:
				currentGuardPosition.Direction = Right
			case Down:
				currentGuardPosition.Direction = Left
			case Right:
				currentGuardPosition.Direction = Down
			case Left:
				currentGuardPosition.Direction = Up
			}
		} else {
			break
		}
	}

	return GuardPosition{Position: Position{X: nextX, Y: nextY}, Direction: currentGuardPosition.Direction}
}

func getNextPosition(position GuardPosition) (int, int) {
	nextX := position.X
	nextY := position.Y

	switch position.Direction {
	case Up:
		nextY--
	case Down:
		nextY++
	case Right:
		nextX++
	case Left:
		nextX--
	}

	return nextX, nextY
}

func getUnique(positions []GuardPosition) []GuardPosition {
	var uniquePositions []GuardPosition

	for _, p := range positions {
		if isUnique(uniquePositions, p) {
			uniquePositions = append(uniquePositions, p)
		}
	}

	return uniquePositions
}

func isUnique(uniquePositions []GuardPosition, p GuardPosition) bool {
	isUnique := true
	for _, u := range uniquePositions {
		if p.X == u.X && p.Y == u.Y {
			isUnique = false
			break
		}
	}

	return isUnique
}
