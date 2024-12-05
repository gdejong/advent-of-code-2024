package math

func Abs(value int) int {
	if value < 0 {
		value = value * -1
	}

	return value
}
