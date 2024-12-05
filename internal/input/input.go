package input

import (
	"github.com/gdejong/advent-of-code-2024/internal/must"
	"strconv"
)

func ToIntegers(input []string) []int {
	integers := make([]int, len(input))

	// Convert to integer values
	for k, v := range input {
		integers[k] = must.NoError(strconv.Atoi(v))
	}

	return integers
}
