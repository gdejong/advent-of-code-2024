package slices

// All accepts a slice and a predicate function. It will return true if the predicate is true for all slice elements.
func All[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}

// CopyAndRemoveIndex returns a new slice with the given i element removed.
func CopyAndRemoveIndex(slice []int, i int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice) // make a copy so we do not alter the original slice

	return append(newSlice[:i], newSlice[i+1:]...)
}
