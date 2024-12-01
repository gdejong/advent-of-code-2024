package main

import (
	"bufio"
	"fmt"
	"github.com/gdejong/advent-of-code-2024/internal"
	"io"
	"os"
	"slices"
)

func main() {
	f := internal.Must(os.Open("cmd/day1/real_input.txt"))

	part1answer := part1(f)
	fmt.Println(part1answer)

	// Reset the file so we can it again for part 2
	internal.Must(f.Seek(0, io.SeekStart))

	part2answer := part2(f)
	fmt.Println(part2answer)
}

func part2(f *os.File) int {
	s := bufio.NewScanner(f)

	var leftList []int
	var rightList []int

	for s.Scan() {
		line := s.Text()

		var leftNumber int
		var rightNumber int
		internal.Must(fmt.Sscanf(line, "%d %d", &leftNumber, &rightNumber))
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

func part1(f *os.File) int {
	s := bufio.NewScanner(f)

	var leftList []int
	var rightList []int

	for s.Scan() {
		line := s.Text()

		var leftNumber int
		var rightNumber int
		internal.Must(fmt.Sscanf(line, "%d %d", &leftNumber, &rightNumber))
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
