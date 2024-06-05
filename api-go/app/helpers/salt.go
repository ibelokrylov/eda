package helpers

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateSalt(length int) (string, error) {
	byteLength := (length*6 + 7) / 8
	bytes := make([]byte, byteLength)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	randomString := base64.URLEncoding.EncodeToString(bytes)

	if len(randomString) > length {
		randomString = randomString[:length]
	}

	return randomString, nil
}
