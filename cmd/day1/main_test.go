package main

import (
	"github.com/gdejong/advent-of-code-2024/internal"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	f := internal.Must(os.Open("test_input.txt"))

	answer := part1(f)

	if answer != 11 {
		t.Errorf("wrong answer, got: %d", answer)
	}
}

func TestPart2(t *testing.T) {
	f := internal.Must(os.Open("test_input.txt"))

	answer := part2(f)

	if answer != 31 {
		t.Errorf("wrong answer, got: %d", answer)
	}
}
