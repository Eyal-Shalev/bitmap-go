package bitmap

import (
	"fmt"
)

type MapFunc[T any] func(idx int, isSet bool) T

func Map[T any](bm BitMap, f MapFunc[T]) []T {
	results := make([]T, bm.Length())
	for idx := 0; idx < bm.Length(); idx++ {
		isSet, err := bm.IsSet(idx)
		if err != nil {
			panic(fmt.Sprintf("bitmap.BitMap.ForEach: unexpected %s", err.Error()))
		}
		results[idx] = f(idx, isSet)
	}
	return results
}

type ReduceFunc[T any] func(idx int, isSet bool, Accumulator T) T

func ReduceWithInit[T any](bm BitMap, f ReduceFunc[T], initial T) T {
	current := initial
	for idx := 0; idx < bm.Length(); idx++ {
		isSet, err := bm.IsSet(idx)
		if err != nil {
			panic(fmt.Sprintf("bitmap.BitMap.ForEach: unexpected %s", err.Error()))
		}
		current = f(idx, isSet, current)
	}
	return current
}

func Reduce[T any](bm BitMap, f ReduceFunc[T]) T {
	var zero T
	return ReduceWithInit(bm, f, zero)
}
