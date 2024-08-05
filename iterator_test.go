//go:build go1.23

package bitmap_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitMap_Iter(t *testing.T) {
	bm := slices.Clone(exampleSmall)
	idx := 0
	for isSet := range bm.Iter() {
		assert.Equal(t, slices.Contains(exampleSmallSetPositions, idx), isSet)
		idx++
	}
}

func TestBitMap_Iter2(t *testing.T) {
	bm := slices.Clone(exampleSmall)
	for idx, isSet := range bm.Iter2() {
		assert.Equal(t, slices.Contains(exampleSmallSetPositions, idx), isSet)
	}
}
