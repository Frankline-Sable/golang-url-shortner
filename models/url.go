package models

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateShortURL creates a unique 6-character short URL
func GenerateShortURL() string {
	randomBytes := make([]byte, 6) // Create a byte slice(array) of length 6

	_, err := rand.Read(randomBytes) // fill with random bytes

	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(randomBytes)[:6] //Encode to base64 string and return first 6 characters

}
