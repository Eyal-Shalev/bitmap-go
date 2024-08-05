//go:build go1.23

package bitmap

import (
	"iter"
)

func (bm *BitMap) Iter() iter.Seq[bool] {
	return func(yield func(bool) bool) {
		for idx := 0; idx < bm.Length(); idx++ {
			isSet, _ := bm.IsSet(idx)
			if !yield(isSet) {
				return
			}
		}
	}
}

func (bm *BitMap) Iter2() iter.Seq2[int, bool] {
	return func(yield func(int, bool) bool) {
		for idx := 0; idx < bm.Length(); idx++ {
			isSet, _ := bm.IsSet(idx)
			if !yield(idx, isSet) {
				return
			}
		}
	}
}
