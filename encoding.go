package bitmap

import (
	"encoding/base64"
	"fmt"
	"slices"
)

func (bm *BitMap) UnmarshalBinary(data []byte) error {
	(*bm).data = slices.Clone(data)
	return nil
}

func (bm *BitMap) MarshalBinary() ([]byte, error) {
	if bm == nil {
		return nil, nil
	}
	return slices.Clone(bm.data), nil
}

func (bm *BitMap) UnmarshalText(data []byte) error {
	result, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return fmt.Errorf("bitmap.BitMap.UnmarshalText: %w", err)
	}
	(*bm).data = result
	return nil
}

func (bm *BitMap) MarshalText() ([]byte, error) {
	if bm == nil {
		return nil, nil
	}
	return []byte(base64.StdEncoding.EncodeToString(bm.data)), nil
}
