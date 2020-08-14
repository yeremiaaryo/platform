package utils

import (
	"crypto/sha1"
	"fmt"
)

func GenerateSHA1(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	bs := h.Sum(nil)
	result := fmt.Sprintf("%x", bs)
	return result
}
