package api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

func ShaHMAC(input, key string) string {
	keySign := []byte(key)
	h := hmac.New(sha1.New, keySign)
	h.Write([]byte(input))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
