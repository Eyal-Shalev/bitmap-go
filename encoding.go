package bitmap

import (
	"encoding/base64"
	"fmt"
	"slices"
)

func (bm *BitMap) UnmarshalBinary(data []byte) error {
	(*bm).padding = int(data[0])
	(*bm).data = slices.Clone(data[1:])
	return nil
}

func (bm *BitMap) MarshalBinary() ([]byte, error) {
	if bm == nil {
		return nil, nil
	}
	return append([]byte{byte(bm.padding)}, bm.data...), nil
}

func (bm *BitMap) UnmarshalText(textData []byte) error {
	data, err := base64.StdEncoding.DecodeString(string(textData))
	if err != nil {
		return fmt.Errorf("bitmap.BitMap.UnmarshalText: %w", err)
	}
	err = bm.UnmarshalBinary(data)
	if err != nil {
		return fmt.Errorf("bitmap.BitMap.UnmarshalText: %w", err)
	}
	return nil
}

func (bm *BitMap) MarshalText() ([]byte, error) {
	binaryData, err := bm.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("bitmap.BitMap.MarshalText: %w", err)
	}
	return []byte(base64.StdEncoding.EncodeToString(binaryData)), nil
}
