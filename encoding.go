package bitmap

import (
	"encoding/base64"
	"fmt"
	"slices"
)

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
