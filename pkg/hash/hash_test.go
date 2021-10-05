package hash_test

import (
	"testing"

	"github.com/nawtinhaz/jwt-auth/pkg/hash"
)

func TestKeygen(t *testing.T) {
	const size = 64
	key, err := hash.Keygen(size)
	if err != nil {
		t.Error("failed to generate hash:", err)
	}

	if got, exp := len(key), size; got != exp {
		t.Errorf("unexpected key size: got %v, expected %v", got, exp)
	}
}

func TestPassword(t *testing.T) {
	const size = 64
	salt, _ := hash.Keygen(size)
	key, err := hash.Password([]byte("password"), salt, size)
	if err != nil {
		t.Error("failed to hash password:", err)
	}

	if got, exp := len(key), size; got != exp {
		t.Errorf("unexpected hash size: got %v, expected %v", got, exp)
	}
}

func TestPasswordCompare(t *testing.T) {
	const size = 64
	salt, _ := hash.Keygen(size)
	key, err := hash.Password([]byte("password"), salt, size)
	if err != nil {
		t.Error("failed to hash password:", err)
	}

	if got, exp := len(key), size; got != exp {
		t.Errorf("unexpected hash size: got %v, expected %v", got, exp)
	}

	tcs := []struct {
		name           string
		password       []byte
		hashedPassword func() []byte
		exp            bool
	}{
		{
			name:     "equal passwords",
			password: []byte("password"),
			hashedPassword: func() []byte {
				key, _ := hash.Password([]byte("password"), salt, size)
				return key
			},
			exp: true,
		},
		{
			name:     "different passwords",
			password: []byte("drowssap"),
			hashedPassword: func() []byte {
				key, _ := hash.Password([]byte("password"), salt, size)
				return key
			},
			exp: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := hash.PasswordCompare(tc.password, salt, size, tc.hashedPassword())
			if err != nil {
				t.Error("unexpected error when comparing hashed passwords:", err)
			}

			if tc.exp != got {
				t.Errorf("unxpected comparison result for %q: got %v, exp %v", tc.name, got, tc.exp)
			}
		})
	}
}
