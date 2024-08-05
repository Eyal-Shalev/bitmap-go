package bitmap

import (
	"fmt"
)

type stringError string

func (e stringError) Error() string {
	return string(e)
}

const Error = stringError("bitmap.Error")

type IndexOutOfBoundError struct {
	Index  int
	Length int
}

func (e *IndexOutOfBoundError) Error() string {
	return fmt.Sprintf("bitmap.IndexOutOfBoundError: index=%d, length=%d", e.Index, e.Length)
}

func (e *IndexOutOfBoundError) Unwrap() error {
	return Error
}

func (e *IndexOutOfBoundError) Is(other error) bool {
	switch other.(type) {
	case *IndexOutOfBoundError:
		return true
	default:
		return other == Error
	}
}

type InvalidLengthError struct {
	Length int
}

func (e *InvalidLengthError) Error() string {
	prefix := fmt.Sprintf("bitmap.InvalidLengthError{Length: %d}: ", e.Length)
	if e.Length < 0 {
		return prefix + "Length must be larger than 0"
	}
	return prefix + "Length must be a multiple of 8"
}

func (e *InvalidLengthError) Unwrap() error {
	return Error
}

func (e *InvalidLengthError) Is(other error) bool {
	switch other.(type) {
	case *InvalidLengthError:
		return true
	default:
		return other == Error
	}
}
