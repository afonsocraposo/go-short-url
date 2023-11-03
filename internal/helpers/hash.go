package helpers

import (
    "encoding/base64"
	"crypto/sha256"
)

func HashString(value string, size int) string {
    if size == 0 {
        size = 16
    }
    h := sha256.New()
    h.Write([]byte(value))
    bs := h.Sum(nil)
    hash := base64.URLEncoding.EncodeToString(bs[:])
    return hash[:size]
}

