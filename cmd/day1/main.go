package main

import (
	"bufio"
	"fmt"
	"github.com/gdejong/advent-of-code-2024/internal"
	"os"
	"slices"
)

func main() {
	f := internal.Must(os.Open("cmd/day1/real_input.txt"))

	part1answer := part1(f)

	fmt.Println(part1answer)
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
