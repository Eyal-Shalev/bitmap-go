package bitmap_test

import (
	"testing"

	"github.com/Eyal-Shalev/bitmap-go/v0"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	indexes := bitmap.Map(exampleSmall, func(idx int, isSet bool) int {
		return idx
	})
	assert.Equal(t, exampleSmall.Length(), len(indexes))
	for idxIdx, idx := range indexes {
		assert.Equal(t, idx, idxIdx)
	}
}

func TestReduce(t *testing.T) {
	isAnySet := bitmap.Reduce(exampleSmall, func(idx int, isSet bool, Accumulator bool) bool {
		return Accumulator || isSet
	})
	assert.True(t, isAnySet)

	areAllSet := bitmap.Reduce(exampleSmall, func(idx int, isSet bool, Accumulator bool) bool {
		return Accumulator && isSet
	})
	assert.False(t, areAllSet)
}

func andFn(_ int, isSet bool, Accumulator bool) bool {
	return isSet && Accumulator
}
func orFn(_ int, isSet bool, Accumulator bool) bool {
	return isSet || Accumulator
}

func TestReduceWithInit(t *testing.T) {
	type args[T any] struct {
		bm      bitmap.BitMap
		f       bitmap.ReduceFunc[T]
		initial bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[bool]{
		{
			name: "reduceWithInit(0, andFn, false)",
			args: args[bool]{
				bm:      bitmap.BitMap{0},
				f:       andFn,
				initial: false,
			},
			want: false,
		},
		{
			name: "reduceWithInit(1, andFn, false)",
			args: args[bool]{
				bm:      bitmap.BitMap{255},
				f:       andFn,
				initial: false,
			},
			want: false,
		},
		{
			name: "reduceWithInit(0, andFn, true)",
			args: args[bool]{
				bm:      bitmap.BitMap{0},
				f:       andFn,
				initial: true,
			},
			want: false,
		},
		{
			name: "reduceWithInit(255, andFn, true)",
			args: args[bool]{
				bm:      bitmap.BitMap{255},
				f:       andFn,
				initial: true,
			},
			want: true,
		},
		{
			name: "reduceWithInit(0, orFn, false)",
			args: args[bool]{
				bm:      bitmap.BitMap{0},
				f:       orFn,
				initial: false,
			},
			want: false,
		},
		{
			name: "reduceWithInit(1, orFn, false)",
			args: args[bool]{
				bm:      bitmap.BitMap{1},
				f:       orFn,
				initial: false,
			},
			want: true,
		},
		{
			name: "reduceWithInit(0, orFn, true)",
			args: args[bool]{
				bm:      bitmap.BitMap{0},
				f:       orFn,
				initial: true,
			},
			want: true,
		},
		{
			name: "reduceWithInit(255, orFn, true)",
			args: args[bool]{
				bm:      bitmap.BitMap{1},
				f:       orFn,
				initial: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, bitmap.ReduceWithInit(tt.args.bm, tt.args.f, tt.args.initial), "ReduceWithInit(%v, %v, %v)", tt.args.bm, tt.args.f, tt.args.initial)
		})
	}
}
