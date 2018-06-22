package util

import (
	"crypto/sha1"
	"fmt"
)

// sha1加密
func Sha1hex(data []byte) string {
	h := sha1.New()
	h.Write(data)
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
