package bitmap

import (
	"fmt"
	"strings"
)

type BitMap struct {
	data []byte
}

func (bm *BitMap) String() string {
	if bm == nil {
		return ""
	}
	sb := &strings.Builder{}
	for _, b := range bm.data {
		_, _ = fmt.Fprintf(sb, "%08b", b)
	}
	return sb.String()
}

func (bm *BitMap) GoString() string {
	if bm == nil {
		return "bitmap.BitMap(nil)"
	}
	sb := &strings.Builder{}
	sb.WriteString("bitmap.BitMap{")
	for idx, b := range bm.data {
		if idx > 0 {
			sb.WriteString(", ")
		}
		_, _ = fmt.Fprintf(sb, "0b%08b", b)
	}
	sb.WriteString("}")
	return sb.String()
}

func (bm *BitMap) calcPosAndMask(idx int) (int, byte, error) {
	if bm == nil {
		return 0, 0, &IndexOutOfBoundError{Index: idx, Length: 0}
	}
	if idx/8 >= len(bm.data) || idx < 0 {
		return 0, 0, &IndexOutOfBoundError{Index: idx, Length: len(bm.data)}
	}
	pos := idx / 8
	bitPos := idx % 8
	shiftBy := 8 - bitPos - 1
	mask := byte(1 << shiftBy)
	return pos, mask, nil
}

func (bm *BitMap) Length() int {
	if bm == nil {
		return 0
	}
	return len(bm.data) * 8
}

func (bm *BitMap) IsSet(idx int) (bool, error) {
	pos, mask, err := bm.calcPosAndMask(idx)
	if err != nil {
		return false, err
	}
	byteAtPos := bm.data[pos]
	return byteAtPos&mask != 0, nil
}

func (bm *BitMap) Set(idx int) error {
	pos, mask, err := bm.calcPosAndMask(idx)
	if err != nil {
		return err
	}
	bm.data[pos] = bm.data[pos] | mask
	return nil
}

func (bm *BitMap) UnSet(idx int) error {
	pos, mask, err := bm.calcPosAndMask(idx)
	if err != nil {
		return err
	}
	bm.data[pos] = bm.data[pos] ^ mask
	return nil
}

func (bm *BitMap) SetVal(idx int, isSet bool) error {
	if isSet {
		return bm.Set(idx)
	} else {
		return bm.UnSet(idx)
	}
}

func New(length int) (*BitMap, error) {
	if length < 0 || length%8 != 0 {
		return nil, &InvalidLengthError{Length: length}
	}
	return &BitMap{data: make([]byte, length/8)}, nil
}

func NewFromBytes(data []byte) *BitMap {
	return &BitMap{data: data}
}
