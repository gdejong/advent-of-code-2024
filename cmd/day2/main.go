package main

import (
	"bufio"
	"fmt"
	"github.com/gdejong/advent-of-code-2024/internal/input"
	"github.com/gdejong/advent-of-code-2024/internal/math"
	"github.com/gdejong/advent-of-code-2024/internal/must"
	"github.com/gdejong/advent-of-code-2024/internal/slices"
	"os"
	"strings"
)

func main() {
	f := must.NoError(os.Open("cmd/day2/real_input.txt"))

	fmt.Println(part1(f))
}

func part1(f *os.File) int {
	s := bufio.NewScanner(f)

	safeCounter := 0
	for s.Scan() {
		line := s.Text()

		levels := input.ToIntegers(strings.Fields(line))

		if isSafeReport(levels) {
			safeCounter++
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
