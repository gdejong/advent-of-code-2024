package main

import (
	"github.com/gdejong/advent-of-code-2024/internal/must"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	f := must.NoError(os.Open("test_input.txt"))

	answer := part1(f)

	if answer != 2 {
		t.Errorf("wrong answer, got: %d", answer)
	}
}
