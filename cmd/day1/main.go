package main

import (
	"fmt"
	"github.com/gdejong/advent-of-code-2024/internal/input"
	"github.com/gdejong/advent-of-code-2024/internal/must"
	"slices"
)

func main() {
	lines := input.Content("cmd/day1/real_input.txt")

	fmt.Println(part1(lines))

	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	var leftList []int
	var rightList []int

	for _, line := range lines {
		var leftNumber int
		var rightNumber int
		must.NoError(fmt.Sscanf(line, "%d %d", &leftNumber, &rightNumber))
		leftList = append(leftList, leftNumber)
		rightList = append(rightList, rightNumber)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	if len(leftList) != len(rightList) {
		panic("expected same length")
	}

	distanceSum := 0
	for i := 0; i < len(leftList); i++ {
		distance := leftList[i] - rightList[i]

		// Take the absolute value
		if distance < 0 {
			distance = distance * -1
		}

		distanceSum += distance
	}

	return distanceSum
}

func part2(lines []string) int {
	var leftList []int
	var rightList []int

	for _, line := range lines {
		var leftNumber int
		var rightNumber int
		must.NoError(fmt.Sscanf(line, "%d %d", &leftNumber, &rightNumber))
		leftList = append(leftList, leftNumber)
		rightList = append(rightList, rightNumber)
	}

	counterMap := make(map[int]int)

	for _, value := range rightList {
		if _, ok := counterMap[value]; ok {
			counterMap[value]++
		} else {
			counterMap[value] = 1
		}
	}

	similarityScore := 0
	for _, value := range leftList {
		counterMapValue, ok := counterMap[value]
		if !ok {
			continue
		}

		similarityScore += value * counterMapValue
	}

	return similarityScore
}
