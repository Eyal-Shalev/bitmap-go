package bitmap

import (
	"fmt"
	"slices"
	"strings"
)

type BitMap struct {
	data    []byte
	padding int
}

func (bm *BitMap) String() string {
	if bm == nil {
		return ""
	}
	sb := &strings.Builder{}
	for idx, b := range bm.data {
		if idx == len(bm.data)-1 && bm.padding > 0 {
			_, _ = fmt.Fprintf(sb, "%0*b", bm.padding, b>>bm.padding)
		} else {
			_, _ = fmt.Fprintf(sb, "%08b", b)
		}
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
		if idx == len(bm.data)-1 && bm.padding > 0 {
			_, _ = fmt.Fprintf(sb, "0b%0*b", bm.padding, b>>bm.padding)
		} else {
			_, _ = fmt.Fprintf(sb, "0b%08b", b)
		}
	}
	sb.WriteString("}")
	return sb.String()
}

func (bm *BitMap) calcPosAndMask(idx int) (int, byte, error) {
	if bm == nil {
		return 0, 0, &IndexOutOfBoundError{Index: idx, Length: 0}
	}
	if idx >= bm.Length() || idx < 0 {
		return 0, 0, &IndexOutOfBoundError{Index: idx, Length: bm.Length()}
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
	return len(bm.data)*8 - bm.padding
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
	if length < 0 {
		return nil, &InvalidLengthError{Length: length}
	}

	padding := (8 - (length % 8)) % 8
	bytesLength := length / 8
	if padding > 0 {
		bytesLength += 1
	}

	return &BitMap{data: make([]byte, bytesLength), padding: padding}, nil
}

func NewFromBytes(data []byte, padding int) (*BitMap, error) {
	if padding < -8 || padding > 8 {
		return nil, &InvalidPaddingError{Padding: padding}
	}
	data = slices.Clone(data)
	if padding > 0 {
		data = append(data, 0)
	} else if padding < 0 {
		padding = -padding
		data[len(data)-1] = (data[len(data)-1] >> padding) << padding
	}
	return &BitMap{data: data, padding: padding}, nil
}
