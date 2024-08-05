package bitmap_test

import (
	"fmt"
	"testing"

	"github.com/Eyal-Shalev/bitmap-go/v1"
	"github.com/stretchr/testify/assert"
)

func TestIndexOutOfBoundError_Unwrap(t *testing.T) {
	err := &bitmap.IndexOutOfBoundError{Index: 2, Length: 1}

	assert.ErrorIs(t, err, bitmap.Error)
	assert.NotErrorIs(t, err, fmt.Errorf("some other error"))
	assert.NotErrorIs(t, fmt.Errorf("some other error"), bitmap.Error)
}
