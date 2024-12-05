package input

import (
	"bufio"
	"github.com/gdejong/advent-of-code-2024/internal/must"
	"os"
	"strconv"
)

func Content(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	lines := make([]string, 0)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func ToIntegers(input []string) []int {
	integers := make([]int, len(input))

	// Convert to integer values
	for k, v := range input {
		integers[k] = must.NoError(strconv.Atoi(v))
	}

	return integers
}
