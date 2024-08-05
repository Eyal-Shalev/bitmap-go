package bitmap_test

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBitMap_BinaryMarshaller(t *testing.T) {
	msg := []byte(`Hello World`)
	var bm bitmap.BitMap
	require.NoError(t, bm.UnmarshalBinary(msg))
	result, err := bm.MarshalBinary()
	require.NoError(t, err)
	assert.Equal(t, msg, result)
}

func TestBitMap_jsonMarshaller(t *testing.T) {
	const msgStr = `Hello World`
	msgBase64 := base64.StdEncoding.EncodeToString([]byte(msgStr))
	msgBM := bitmap.BitMap(msgStr)
	jsonStr, err := json.Marshal(msgBM)
	require.NoError(t, err)
	assert.Equal(t, jsonStr, []byte(`"`+msgBase64+`"`))
	var result bitmap.BitMap
	require.NoError(t, json.Unmarshal(jsonStr, &result))
	assert.Equal(t, msgBM, result)
}
