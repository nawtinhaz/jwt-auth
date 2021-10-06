// Package sign provides message signing and validation functionality.
package sign

import (
	"crypto"
	"crypto/hmac"
)

// SHA256 is the hashing algorithm used for signing.
const SHA256 = crypto.SHA256

// Sign signs msg with key using hmac.
func Sign(msg, key []byte) []byte {
	mac := hmac.New(SHA256.New, key)
	return mac.Sum(msg)
}

// Valid verifies if msg produces sig after signing.
func Valid(msg, sig, key []byte) bool {
	mac := hmac.New(SHA256.New, key)
	digest := mac.Sum(msg)
	return hmac.Equal(sig, digest)
}
