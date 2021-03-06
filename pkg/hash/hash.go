// Package hash implements key generation and password hashing operations.
package hash

import (
	"crypto/rand"
	"crypto/subtle"

	"golang.org/x/crypto/scrypt"
)

const DigestSize = 64

// Keygen generates a cryptographically secure key of given size.
func Keygen() ([]byte, error) {
	key := make([]byte, DigestSize)

	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

// Password hashes a given password with a given salt.
// It returns an hash of the given size.
func Password(pass, salt []byte) ([]byte, error) {
	result, err := scrypt.Key(pass, salt, 1<<15, 8, 1, DigestSize)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PasswordCompare allows comparison between a given password and a given hashed password.
//
// Returns true if after hashing the given password with the given salt the output
// hash of the given size is equal to the given hashed password. Else returns false.
func PasswordCompare(pass, salt, hashedPass []byte) error {
	result, err := scrypt.Key(pass, salt, 1<<15, 8, 1, DigestSize)
	if err != nil {
		return err
	}
	if subtle.ConstantTimeCompare(result, hashedPass) != 1 {
		return ErrMismatchedHashandPassword
	}
	return nil
}
