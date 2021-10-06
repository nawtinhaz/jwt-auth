// Package hash implements key generation and password hashing operations.
package hash

import (
	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/scrypt"
)

const digestSize = 64

// Keygen generates a cryptographically secure key of given size.
func Keygen() ([]byte, error) {
	key := make([]byte, digestSize)

	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

// Password hashes a given password with a given salt.
// It returns an hash of the given size.
func Password(pass, salt []byte) ([]byte, error) {
	result, err := scrypt.Key(pass, salt, 1<<15, 8, 1, digestSize)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PasswordCompare allows comparison between a given password and a given hashed password.
//
// Returns true if after hashing the given password with the given salt the output
// hash of the given size is equal to the given hashed password. Else returns false.
func PasswordCompare(pass, salt, hashedPass []byte) (bool, error) {
	result, err := scrypt.Key(pass, salt, 1<<15, 8, 1, digestSize)
	if err != nil {
		return false, err
	}
	return bytes.Equal(result, hashedPass), nil
}
