package bitmap

import (
	"slices"
)

func (bm *BitMap) Clone() *BitMap {
	return &BitMap{data: slices.Clone(bm.data)}
}
