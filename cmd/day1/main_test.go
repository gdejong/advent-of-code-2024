package main

import (
	"github.com/gdejong/advent-of-code-2024/internal/input"
	"testing"
)

func TestPart1(t *testing.T) {
	lines := input.Content("test_input.txt")

	answer := part1(lines)

	if answer != 11 {
		t.Errorf("wrong answer, got: %d", answer)
	}
}

func TestPart2(t *testing.T) {
	lines := input.Content("test_input.txt")

	answer := part2(lines)

	if answer != 31 {
		t.Errorf("wrong answer, got: %d", answer)
	}
}
