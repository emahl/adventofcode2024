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

	var visitedPositions []GuardPosition
	visitedPositions = append(visitedPositions, guardStartingPosition)
	for {
		newGuardPosition := moveGuard(mappedArea, visitedPositions[len(visitedPositions)-1])
		if newGuardPosition.X == -1 && newGuardPosition.Y == -1 {
			break
		}
		visitedPositions = append(visitedPositions, newGuardPosition)
	}

	uniquePositions := unique(visitedPositions)
	sum := len(uniquePositions)
	fmt.Println("The number of distinct guard positions:", sum)
}

func unique(positions []GuardPosition) []GuardPosition {
	var uniquePositions []GuardPosition

	for _, p := range positions {
		isUnique := true
		for _, u := range uniquePositions {
			if p.X == u.X && p.Y == u.Y {
				isUnique = false
				break
			}
		}
		if isUnique {
			uniquePositions = append(uniquePositions, p)
		}
	}

	return uniquePositions
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
