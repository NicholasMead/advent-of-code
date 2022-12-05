package common

func ToSlice[T any](input <-chan T) <-chan []T {
	output := make(chan []T, 1)

	go func(in <-chan T, out chan<- []T) {
		defer close(output)
		slice := []T{}
		for message := range in {
			slice = append(slice, message)
		}
		out <- slice
	}(input, output)

	return output
}
