package main

import (
	"bufio"
	"fmt"
	"github.com/gdejong/advent-of-code-2024/internal/input"
	"github.com/gdejong/advent-of-code-2024/internal/math"
	"github.com/gdejong/advent-of-code-2024/internal/must"
	"github.com/gdejong/advent-of-code-2024/internal/slices"
	"io"
	"os"
	"strings"
)

func main() {
	f := must.NoError(os.Open("cmd/day2/real_input.txt"))

	fmt.Println(part1(f))

	// Reset the file so we can it again for part 2
	must.NoError(f.Seek(0, io.SeekStart))

	fmt.Println(part2(f))
}

func part1(f *os.File) int {
	s := bufio.NewScanner(f)

	safeCounter := 0
	for s.Scan() {
		line := s.Text()

		report := input.ToIntegers(strings.Fields(line))

		if isSafeReport(report) {
			safeCounter++
		}
	}

	return safeCounter
}

func part2(f *os.File) int {
	s := bufio.NewScanner(f)

	safeCounter := 0
	for s.Scan() {
		line := s.Text()

		report := input.ToIntegers(strings.Fields(line))

		// Create all possible options by continually removing one level
		var reportsToCheck [][]int
		reportsToCheck = append(reportsToCheck, report)
		for index := range len(report) {
			reportsToCheck = append(reportsToCheck, remove(report, index))
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

func remove(slice []int, s int) []int {
	s2 := make([]int, len(slice))
	copy(s2, slice)
	return append(s2[:s], s2[s+1:]...)
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
