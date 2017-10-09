package challenge1

import (
	"encoding/base64"
	"encoding/hex"
)

// DecodeHex accepts a hex-encoded string and returns a decoded []byte
func DecodeHex(value string) []byte {
	decoded, err := hex.DecodeString(value)
	if err != nil {
		panic(err)
	}
	return decoded
}

// EncodeBase64 accepts a base64 encoded []bytes and returns a decoded string
func EncodeBase64(value []byte) string {
	return base64.URLEncoding.EncodeToString(value)
}
