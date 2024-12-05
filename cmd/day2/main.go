package main

import (
	"fmt"
	"github.com/gdejong/advent-of-code-2024/internal/input"
	"github.com/gdejong/advent-of-code-2024/internal/math"
	"github.com/gdejong/advent-of-code-2024/internal/slices"
	"strings"
)

func main() {
	lines := input.Content("cmd/day2/real_input.txt")

	fmt.Println(part1(lines))

	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	safeCounter := 0
	for _, line := range lines {
		report := input.ToIntegers(strings.Fields(line))

		if isSafeReport(report) {
			safeCounter++
		}
	}

	return safeCounter
}

func part2(lines []string) int {
	safeCounter := 0
	for _, line := range lines {
		report := input.ToIntegers(strings.Fields(line))

		// Create all possible options by continually removing one level
		var reportsToCheck [][]int
		reportsToCheck = append(reportsToCheck, report)
		for index := range len(report) {
			reportsToCheck = append(reportsToCheck, slices.CopyAndRemoveIndex(report, index))
		}

		for _, r := range reportsToCheck {
			if isSafeReport(r) {
				safeCounter++
				break
			}
		}
	}

	return safeCounter
}

func isSafeReport(levels []int) bool {
	// Calculate the differences
	diffs := make([]int, len(levels)-1)
	for index := range len(levels) - 1 {
		diffs[index] = levels[index+1] - levels[index]
	}

	allIncreasing := func(val int) bool { return val > 0 }
	allDecreasing := func(val int) bool { return val < 0 }
	allBetweenOneAndThree := func(val int) bool { return math.Abs(val) >= 1 && math.Abs(val) <= 3 }

	isAllIncreasing := slices.All(diffs, allIncreasing)
	isAllDecreasing := slices.All(diffs, allDecreasing)
	IsAllDifferByAtLeastOneAndAtMostThree := slices.All(diffs, allBetweenOneAndThree)

	if (isAllIncreasing || isAllDecreasing) && IsAllDifferByAtLeastOneAndAtMostThree {
		return true
	}

	return false
}
