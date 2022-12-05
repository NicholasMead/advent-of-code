package common

func CopySlice[T any](a []T) []T {
	b := make([]T, len(a))

	for i, v := range a {
		b[i] = v
	}


	return b
}
