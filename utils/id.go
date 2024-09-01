package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateSlug() (string, error) {
    bytes := make([]byte, 6) 
    _, err := rand.Read(bytes)
    if err != nil {
        return "", err
    }
    return base64.RawURLEncoding.EncodeToString(bytes)[:8], nil
}
