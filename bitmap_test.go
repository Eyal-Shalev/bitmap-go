package bitmap_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/Eyal-Shalev/bitmap-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleSmall = bitmap.BitMap{0b00000000, 0b10000000, 0b01000000, 0b00100000, 0b00010000, 0b00001000, 0b00000100, 0b00000010, 0b00000001}
var exampleSmallSetPositions = []int{8 * 1, 8*2 + 1, 8*3 + 2, 8*4 + 3, 8*5 + 4, 8*6 + 5, 8*7 + 6, 8*8 + 7}

var exampleRandom = bitmap.BitMap{
	0b00101010,
	0b01000101,
}
var exampleRandomSetPositions = []int{2, 4, 6, 8*1 + 1, 8*1 + 5, 8*1 + 7}

func ExampleBitMap_String() {
	fmt.Print(exampleSmall)
	// Output: 000000001000000001000000001000000001000000001000000001000000001000000001
}

func ExampleBitMap_GoString() {
	fmt.Printf("%#v", exampleSmall)
	// Output: bitmap.BitMap{0b00000000, 0b10000000, 0b01000000, 0b00100000, 0b00010000, 0b00001000, 0b00000100, 0b00000010, 0b00000001}
}

func TestBitMap_IsSet(t *testing.T) {
	for idx := 0; idx < exampleSmall.Length(); idx++ {
		t.Run(fmt.Sprintf("exampleSmall[%d]", idx), func(tt *testing.T) {
			isSet, err := exampleSmall.IsSet(idx)
			require.NoError(tt, err)
			assert.Equal(tt, slices.Contains(exampleSmallSetPositions, idx), isSet)
		})
	}
	for idx := 0; idx < exampleRandom.Length(); idx++ {
		t.Run(fmt.Sprintf("exampleRandom[%d]", idx), func(tt *testing.T) {
			isSet, err := exampleRandom.IsSet(idx)
			require.NoError(tt, err)
			assert.Equal(tt, slices.Contains(exampleRandomSetPositions, idx), isSet)
		})
	}
}

func TestBitMap_Set(t *testing.T) {
	bm, err := bitmap.New(16)
	require.NoError(t, err)

	assert.ErrorIs(t, bm.Set(16), &bitmap.IndexOutOfBoundError{})

	for _, idx := range exampleRandomSetPositions {
		require.NoError(t, bm.Set(idx))
	}

	assert.EqualValues(t, exampleRandom, bm)
}

func TestBitMap_UnSet(t *testing.T) {
	bm := slices.Clone(exampleRandom)
	for _, idx := range exampleRandomSetPositions {
		require.NoError(t, bm.UnSet(idx))
	}
	assert.EqualValues(t, bitmap.BitMap{0b00000000, 0b00000000}, bm)
}

func TestBitMap_SetVal(t *testing.T) {
	bm, err := bitmap.New(16)
	require.NoError(t, err)

	assert.ErrorIs(t, bm.SetVal(16, true), &bitmap.IndexOutOfBoundError{})
	assert.ErrorIs(t, bm.SetVal(16, false), &bitmap.IndexOutOfBoundError{})

	for _, idx := range exampleRandomSetPositions {
		require.NoError(t, bm.SetVal(idx, true))
	}

	assert.EqualValues(t, exampleRandom, bm)

	for _, idx := range exampleRandomSetPositions {
		require.NoError(t, bm.SetVal(idx, false))
	}
	assert.EqualValues(t, bitmap.BitMap{0b00000000, 0b00000000}, bm)
}

func TestBitMap_Length(t *testing.T) {
	assert.Equal(t, 0, bitmap.BitMap(nil).Length())
	assert.Equal(t, 8*2, exampleRandom.Length())
	assert.Equal(t, 8*9, exampleSmall.Length())
}
