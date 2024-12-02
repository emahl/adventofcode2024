package day2

import (
	"fmt"
	"math"
	"strings"

	"github.com/emahl/adventofcode2024/shared"
)

type Options struct {
	UseDapemener bool
}

func Run() {
	reports := readReportsFromFile()

	part1(reports)
	part2(reports)
}

func part1(reports [][]int) {
	safeReports := getSafeReports(reports, Options{})
	fmt.Println("Number of safe reports:", len(safeReports))
}

func part2(reports [][]int) {
	safeReports := getSafeReports(reports, Options{UseDapemener: true})
	fmt.Println("Number of safe reports w/ Problem Dampener:", len(safeReports))
}

func readReportsFromFile() [][]int {
	file, scanner := shared.ReadFile("day2/input.txt")
	defer file.Close()

	var reports [][]int

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		var report []int

		for _, s := range split {
			report = append(report, shared.ConvertToNumber(s))
		}
		reports = append(reports, report)
	}

	return reports
}

func getSafeReports(reports [][]int, options Options) [][]int {
	var safeReports [][]int

	for i := range reports {
		if isSafeReport(reports[i]) {
			safeReports = append(safeReports, reports[i])
		} else if options.UseDapemener {
			// Go through and test other alternatives of the report with the j-th element removed
			for j := range reports[i] {
				adjustedReport := make([]int, len(reports[i]))
				copy(adjustedReport, reports[i])
				adjustedReport = append(adjustedReport[:j], adjustedReport[j+1:]...)

				if isSafeReport(adjustedReport) {
					safeReports = append(safeReports, reports[i])
					break
				}
			}
		}
	}

	return safeReports
}

func isSafeReport(report []int) bool {
	shouldIncrease := report[0] < report[1]
	for i := 1; i < len(report); i++ {
		if !isIncreasingOrDecreasing(report[i-1], report[i], shouldIncrease) || !isValidDifference(report[i-1], report[i]) {
			return false
		}
	}

	return true
}

func isValidDifference(a int, b int) bool {
	difference := math.Abs(float64(a - b))
	return difference > 0 && difference < 4
}

func isIncreasingOrDecreasing(a int, b int, shouldIncrease bool) bool {
	if a == b {
		return false
	}

	if shouldIncrease {
		return a < b
	} else {
		return a > b
	}
}
