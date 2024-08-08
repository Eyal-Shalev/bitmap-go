package bitmap_test

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/Eyal-Shalev/bitmap-go"
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
	expectedMsgBytes := []byte{5}
	expectedMsgBytes = append(expectedMsgBytes, []byte(msgStr)...)
	expectedMsgBytes = append(expectedMsgBytes, 0)
	msgBase64 := base64.StdEncoding.EncodeToString(expectedMsgBytes)
	msgBM, err := bitmap.NewFromBytes([]byte(msgStr), 5)
	require.NoError(t, err)
	jsonStr, err := json.Marshal(msgBM)
	require.NoError(t, err)
	assert.Equal(t, string(jsonStr), `"`+msgBase64+`"`)
	var result bitmap.BitMap
	require.NoError(t, json.Unmarshal(jsonStr, &result))
	assert.Equal(t, msgBM, &result)
}
