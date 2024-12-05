package must

// NoError takes in some value and an error, and panics if the error is not nil.
func NoError[t any](something t, err error) t {
	if err != nil {
		panic(err)
	}

	return something
}
