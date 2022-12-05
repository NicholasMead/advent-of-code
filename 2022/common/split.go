package common

func Split[T any](input <-chan T, count int) []chan T {
	outputs := []chan T{}

	for i := 0; i < count; i++ {
		outputs = append(outputs, make(chan T))
	}

	go func() {
		for item := range input {
			for _, output := range outputs {
				output <- item
			}
		}

		for _, output := range outputs {
			close(output)
		}
	}()

	return outputs
}