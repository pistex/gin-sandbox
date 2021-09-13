package helpers

import (
	"crypto"
	"encoding/base64"
	"fmt"
)

func HashString(plaintext string, alg crypto.Hash) string {
	h := alg.New()
	h.Write([]byte(plaintext))
	b := h.Sum(nil)
	return fmt.Sprintf("%x", b)
}

func HashStringToBase64(plaintext string, alg crypto.Hash) string {
	h := alg.New()
	h.Write([]byte(plaintext))
	b := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(b[:])
}
