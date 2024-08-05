package bitmap

import (
	"encoding/base64"
	"fmt"
	"slices"
	"strings"
)

type BitMap []byte

func (bm *BitMap) UnmarshalBinary(data []byte) error {
	*bm = slices.Clone(data)
	return nil
}

func (bm BitMap) MarshalBinary() ([]byte, error) {
	return slices.Clone(bm), nil
}

func (bm *BitMap) UnmarshalText(data []byte) error {
	result, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return fmt.Errorf("bitmap.BitMap.UnmarshalText: %w", err)
	}
	*bm = result
	return nil
}

func (bm BitMap) MarshalText() ([]byte, error) {
	return []byte(base64.StdEncoding.EncodeToString(bm)), nil
}

func (bm BitMap) String() string {
	sb := &strings.Builder{}
	for _, b := range bm {
		_, _ = fmt.Fprintf(sb, "%08b", b)
	}
	return sb.String()
}

func (bm BitMap) GoString() string {
	sb := &strings.Builder{}
	sb.WriteString("bitmap.BitMap{")
	for idx, b := range bm {
		if idx > 0 {
			sb.WriteString(", ")
		}
		_, _ = fmt.Fprintf(sb, "0b%08b", b)
	}
	sb.WriteString("}")
	return sb.String()
}

func (bm BitMap) calcPosAndMask(idx int) (int, byte, error) {
	if idx/8 >= len(bm) || idx < 0 {
		return 0, 0, &IndexOutOfBoundError{Index: idx, Length: len(bm)}
	}
	pos := idx / 8
	bitPos := idx % 8
	shiftBy := 8 - bitPos - 1
	mask := byte(1 << shiftBy)
	return pos, mask, nil
}

func (bm BitMap) Length() int {
	return len(bm) * 8
}

func (bm BitMap) IsSet(idx int) (bool, error) {
	pos, mask, err := bm.calcPosAndMask(idx)
	if err != nil {
		return false, err
	}
	byteAtPos := bm[pos]
	return byteAtPos&mask != 0, nil
}

func (bm BitMap) Set(idx int) error {
	pos, mask, err := bm.calcPosAndMask(idx)
	if err != nil {
		return err
	}
	bm[pos] = bm[pos] | mask
	return nil
}

func (bm BitMap) UnSet(idx int) error {
	pos, mask, err := bm.calcPosAndMask(idx)
	if err != nil {
		return err
	}
	bm[pos] = bm[pos] ^ mask
	return nil
}

func (bm BitMap) SetVal(idx int, isSet bool) error {
	if isSet {
		return bm.Set(idx)
	} else {
		return bm.UnSet(idx)
	}
}

func New(length int) (BitMap, error) {
	if length < 0 || length%8 != 0 {
		return nil, &InvalidLengthError{Length: length}
	}
	return make(BitMap, length/8), nil
}
