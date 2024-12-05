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
