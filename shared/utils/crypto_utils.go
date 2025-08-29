// utils package contains common utility functions that are used across multiple packages in the application
package utils

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// GenerateRandomSecret generates a cryptographically secure random string,
// suitable for use as OAuth 'state' parameters, API secrets, etc.
//
// Returns:
//   - string: A URL-safe base64 encoded random secret of 32 bytes (typically ~43 characters).
//   - error: Any error that may have occurred.
//
// Example:
//   secret := generateRandomSecret()
func GenerateRandomSecret() (string, error) {
	b := make([]byte, 32) 
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(b), nil
}

// GenerateRandomID generates a cryptographically secure random string,
//
// Returns:
//   - string: A a string of length 'length' containing only the letters 'a'-'z', 'A'-'Z', and '0'-'9'.
//	 - error: Any error that may have occurred.
//
func GenerateRandomID(length int) (string, error) {
	id := make([]byte, length)
	for i := range id {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		id[i] = letters[n.Int64()]
	}
	return string(id), nil
}