package sign

import (
	"crypto"
	"crypto/hmac"
)

const SHA256 = crypto.SHA256

func Sign(message, key []byte) []byte {
	mac := hmac.New(SHA256.New, key)
	return mac.Sum(message)
}

func Valid(message, messageSigned, key []byte) bool {
	mac := hmac.New(SHA256.New, key)
	digest := mac.Sum(message)
	return hmac.Equal(messageSigned, digest)
}
