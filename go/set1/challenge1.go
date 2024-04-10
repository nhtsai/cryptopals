package set1

import (
	"encoding/base64"
	"encoding/hex"
)

// HexToBase64 converts a hex string to a base64 string.
func HexToBase64(s string) (string, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
