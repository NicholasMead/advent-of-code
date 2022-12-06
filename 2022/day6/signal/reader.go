package signal

import "errors"

type Reader interface {
	FindStart(<-chan byte) (int, error)
}

type MarkerNotFound struct{}

type reader struct {
	markerLength int
}

func CreateReader(markerLength int) Reader {
	return reader{
		markerLength: markerLength,
	}
}

func (b reader) FindStart(input <-chan byte) (int, error) {
	candidate := make([]byte, 0, b.markerLength)
	position := 0

	for i := range input {
		position++
		candidate = append(candidate, i)

		if len(candidate) == b.markerLength {
			if areUnique(candidate) {
				return position, nil
			} else {
				candidate = candidate[1:]
			}
		}
	}

	return -1, errors.New("Marker not found")
}

func areUnique(a []byte) bool {
	for i, b := range a {
		if contains(a[i+1:], b) {
			return false
		}
	}
	return true
}

func contains(a []byte, b byte) bool {
	if len(a) <= 0 {
		return false
	}

	for _, a := range a {
		if a == b {
			return true
		}
	}
	return false
}
