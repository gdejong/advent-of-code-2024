package main

import (
	"bufio"
	"fmt"
	"github.com/gdejong/advent-of-code-2024/internal"
	"os"
	"strconv"
	"strings"
)

func main() {
	f := internal.Must(os.Open("cmd/day2/real_input.txt"))

	fmt.Println(part1(f))
}

func part1(f *os.File) int {
	s := bufio.NewScanner(f)

	safeCounter := 0
	for s.Scan() {
		line := s.Text()

		fields := strings.Fields(line)
		levels := make([]int, len(fields))

		// Convert to integer values
		for k, v := range fields {
			levels[k] = internal.Must(strconv.Atoi(v))
		}

		// Calculate the differences
		diffs := make([]int, len(levels)-1)
		for index := range len(levels) - 1 {
			diffs[index] = levels[index+1] - levels[index]
		}

		allIncreasing := func(val int) bool { return val > 0 }
		allDecreasing := func(val int) bool { return val < 0 }
		allBetweenOneAndThree := func(val int) bool { return Abs(val) >= 1 && Abs(val) <= 3 }

		isAllIncreasing := All(diffs, allIncreasing)
		isAllDecreasing := All(diffs, allDecreasing)
		IsAllDifferByAtLeastOneAndAtMostThree := All(diffs, allBetweenOneAndThree)

		if (isAllIncreasing || isAllDecreasing) && IsAllDifferByAtLeastOneAndAtMostThree {
			safeCounter++
		}

	}

	return safeCounter
}

func Abs(value int) int {
	if value < 0 {
		value = value * -1
	}

	return value
}

func All[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}
