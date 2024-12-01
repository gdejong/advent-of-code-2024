package main

import (
	"github.com/gdejong/advent-of-code-2024/internal"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	f := internal.Must(os.Open("test_input.txt"))

	part1answer := part1(f)

	if part1answer != 11 {
		t.Errorf("wrong answer, got: %d", part1answer)
	}
}
