package math

// Abs returns the absolute value of value.
func Abs(value int) int {
	if value < 0 {
		value = value * -1
	}

	return value
}
